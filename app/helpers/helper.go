package helpers

import (
	"github.com/gin-gonic/gin"
	"io"
	"os"
)

type Data struct {
	Data interface{} `json:"data"`
}
type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    Data        `json:"data"`
}

func ApiResponse(data interface{}, message string, status int) Response {
	res := Response{Status: status,
		Message: message,
		Errors:  nil,
		Data: Data{
			Data: data,
		},
	}

	return res

}

func ApiResponseError(data interface{}, err error, status int) Response {
	res := Response{
		Status:  status,
		Message: "error",
		Errors:  err.Error(),
		Data: Data{
			Data: data,
		},
	}

	return res
}
func LoggerFile() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}
