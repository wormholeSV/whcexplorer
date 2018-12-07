package routers

import (
	"bytes"
	"fmt"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var supportTxType = map[string]int{
	"0":  0,
	"1":  1,
	"3":  3,
	"4":  4,
	"50": 50,
	"51": 51,
	"53": 53,
	"54": 54,
	"55": 55,
	"56": 56,
	"68": 68,
	"70": 70,
}


type ResultWriter struct {
	resp *bytes.Buffer
	gin.ResponseWriter
}

func (rw *ResultWriter) Write(p []byte) (int, error) {
	size, err := rw.resp.Write(p)
	if err != nil {
		return size, err
	}
	return rw.ResponseWriter.Write(p)
}

func logContext() gin.HandlerFunc {
	return func(c *gin.Context) {

		rw := &ResultWriter{
			resp:           bytes.NewBuffer(nil),
			ResponseWriter: c.Writer,
		}
		c.Writer = rw

		c.Next()

		if c.Request.URL.Path != "/static" {
			logrus.WithFields(logrus.Fields{
				"query":  c.Request.URL,
				"params": c.Request.Form,
			}).Info("request information:")

			logrus.WithFields(logrus.Fields{
				"response": rw.resp.String(),
			}).Info("response information:")
		}
	}
}

func ginLogger() gin.HandlerFunc {
	// not log the followings router
	notlogged := []string{"/static"}

	out := gin.DefaultWriter
	var skip map[string]struct{}

	if length := len(notlogged); length > 0 {
		skip = make(map[string]struct{}, length)

		for _, path := range notlogged {
			skip[path] = struct{}{}
		}
	}

	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		if _, ok := skip[path]; !ok {
			// Stop timer
			end := time.Now()
			latency := end.Sub(start)

			clientIP := c.ClientIP()
			method := c.Request.Method
			statusCode := c.Writer.Status()
			var statusColor, methodColor, resetColor string
			comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

			if raw != "" {
				path = path + "?" + raw
			}

			fmt.Fprintf(out, "[GIN] %v |%s %3d %s| %13v | %15s |%s %-7s %s %s\n%s",
				end.Format("2006/01/02 - 15:04:05"),
				statusColor, statusCode, resetColor,
				latency,
				clientIP,
				methodColor, method, resetColor,
				path,
				comment,
			)
		}
	}
}
