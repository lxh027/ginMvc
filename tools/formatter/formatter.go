package formatter

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

type ReturnType struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

func ApiReturn(status int, msg string, data interface{}) gin.H {
	return gin.H{
		"status":  status,
		"message": msg,
		"data":    data,
	}
}

func BackendApiReturn(status int, msg string, data interface{}) gin.H {
	return gin.H{
		"status": status,
		"msg":    msg,
		"data":   data,
	}
}

func GetLogFormat(param gin.LogFormatterParams) string {
	return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
		param.ClientIP,
		param.TimeStamp.Format(time.RFC1123),
		param.Method,
		param.Path,
		param.Request.Proto,
		param.StatusCode,
		param.Latency,
		param.Request.UserAgent(),
		param.ErrorMessage,
	)
}
