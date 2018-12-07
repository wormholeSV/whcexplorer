package api_test

import (
	"testing"
	"database/sql"
	"github.com/copernet/whcexplorer/util"
	"github.com/gin-gonic/gin"
	model2 "github.com/copernet/whcexplorer/model"
	"github.com/stretchr/testify/assert"
)
var dbDocker *sql.DB
var testRouter *gin.Engine


type PropertyResponse struct {
	Code int
	Message string
	Result struct{
		List []model2.PropertyObj
		PageNo int
		PageSize int
		Total int
	}
}
func TestProperties(t *testing.T) {
	uri := "/explorer/properties"
	var params = []struct{
		Name string
		Input map[string]string
		Code int
		Total int
	}{
		{"test1",map[string]string{"page":"1"}, 0, 0},
	}
	for _, p := range params {
		t.Run(p.Name, func(t *testing.T) {
			body := util.Get(uri, p.Input, testRouter, t)
			res := &PropertyResponse{}
			util.UnMashallJson(body, res, t)
			assert.Equal(t, p.Code, res.Code, "code not equal")
			assert.Equal(t, p.Total, res.Result.Total, "total not equal")
		})
	}
}

func TestProperties2(t *testing.T) {
	assert.Equal(t, 1,1 )
}
