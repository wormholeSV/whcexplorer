package util

import (
	"github.com/gin-gonic/gin"
	"net/http/httptest"
	"bytes"
	"io/ioutil"
	"encoding/json"
	"testing"
	"path/filepath"
	"github.com/ory/dockertest"
	"log"
	"fmt"
	"database/sql"
	"strconv"
	"github.com/copernet/whccommon/model"
	"os"
)

func PostJson(uri string, param map[string]interface{}, router *gin.Engine, t *testing.T) []byte {
	jsonByte,_ := json.Marshal(param)
	req := httptest.NewRequest("POST", uri, bytes.NewReader(jsonByte))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body,err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error("gin unit test PostJson io error")
	}
	return body
}

func PostForm(uri string, param map[string]string, router *gin.Engine, t *testing.T) []byte {
	req := httptest.NewRequest("POST", uri+ParseToStr(param), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error("gin unit test PostForm io error")
	}
	return body
}

func Get(uri string, param map[string]string,router *gin.Engine, t *testing.T) []byte {
	req := httptest.NewRequest("GET", uri+ParseToStr(param), nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	result := w.Result()
	defer result.Body.Close()
	body, err := ioutil.ReadAll(result.Body)
	if err != nil {
		t.Error("gin unit test Get io error")
	}
	return body
}

func ParseToStr(mp map[string]string) string {
	if len(mp) == 0 {
		return ""
	}
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + val
	}
	temp := values[1:]
	values = "?" + temp
	return values
}


func UnMashallJson(body []byte, res interface{}, t *testing.T) {
	if err := json.Unmarshal(body, res); err != nil {
		t.Error("json unmashall error")
	}
}

func SetupDockerDB(dbDocker *sql.DB, relativePath string) (*dockertest.Resource, *dockertest.Pool, *model.DBOption) {
	// uses a sensible default on windows (tcp/http) and linux/osx (socket)
	pool, err := dockertest.NewPool("")
	if err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}

	// pulls an image, creates a container based on it and runs it
	//path := filepath.Join("testdata","schema.sql")
	p, _ := filepath.Abs(relativePath)
	p = filepath.Join(p, "testdata")
	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "mysql",
		Tag: "5.7",
		Env: []string{"MYSQL_ROOT_PASSWORD=secret"},
		Mounts:[]string{p+":/docker-entrypoint-initdb.d/"},
	})

	if err != nil {
		log.Fatalf("Could not start resource: %s", err)
	}

	// exponential backoff-retry, because the application in the container might not be ready to accept connections yet
	if err := pool.Retry(func() error {
		var err error
		url := fmt.Sprintf("root:secret@(localhost:%s)/mysql", resource.GetPort("3306/tcp"))
		//log.Println(url)
		dbDocker, err = sql.Open("mysql", url)
		if err != nil {
			return err
		}
		return dbDocker.Ping()
	}); err != nil {
		log.Fatalf("Could not connect to docker: %s", err)
	}
	port, _ := strconv.Atoi(resource.GetPort("3306/tcp"))
	db := model.DBOption{
		User:"root",
		Passwd:"secret",
		Host:"localhost",
		Port:port,
		Database:"wormhole",
		Log: true,
	}
	return resource, pool, &db
}

func TeardownDockerDB(resource *dockertest.Resource, pool *dockertest.Pool, code int) {
	// You can't defer this because os.Exit doesn't care for defer
	if err := pool.Purge(resource); err != nil {
		log.Fatalf("Could not purge resource: %s", err)
	}
	os.Exit(code)
}