package helpers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"strconv"
	"time"
)

//	type PaginateResponse struct {
//		Response Response
//	}
type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}
type Meta struct {
	CurrentPage int32 `json:"current_page"`
	LastPage    int32 `json:"last_page"`
	PerPage     int32 `json:"per_page"`
	Total       int32 `json:"total"`
}
type Data struct {
	Data  interface{} `json:"data"`
	Links Links       `json:"links"`
	Meta  Meta        `json:"meta"`
}

type ResponsePaginate struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
	Data    interface{} `json:"data"`
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
func ApiResponsePaginate(data interface{}, message string, status int) ResponsePaginate {
	res := ResponsePaginate{Status: status,
		Message: message,
		Errors:  nil,
		Data:    data,
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
	f, err := os.Create("gin.log")
	if err != nil {
		return
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	defer f.Close()
	return
}

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			params.ClientIP,
			params.TimeStamp.Format(time.RFC3339),
			params.Method,
			params.Path,
			params.StatusCode,
			params.Latency,
		)
	})
}

func Paginate(data []interface{}, limit int, page int, host string) Data {
	var total = len(data)
	lastPage := total / limit
	first := host + "?page=1"
	fmt.Printf("%s", lastPage)
	last := host + "?page=" + strconv.Itoa(lastPage)
	prev := ""
	next := host + "?page=" + strconv.Itoa(page+1)
	if page == lastPage {
		next = ""
	}
	if page != 1 {
		prev = host + "?page=" + strconv.Itoa(page-1)
	}
	if page > lastPage {
		res := Data{
			Data: nil,
			Links: Links{
				First: first,
				Last:  last,
				Prev:  prev,
				Next:  "",
			},
			Meta: Meta{
				CurrentPage: int32(page),
				LastPage:    int32(lastPage),
				PerPage:     int32(limit),
				Total:       int32(total),
			},
		}
		return res
	}

	start := (page - 1) * limit
	end := limit * page
	data = data[start:end]

	res := Data{
		Data: data,
		Links: Links{
			First: first,
			Last:  last,
			Prev:  prev,
			Next:  next,
		},
		Meta: Meta{
			CurrentPage: int32(page),
			LastPage:    int32(lastPage),
			PerPage:     int32(limit),
			Total:       int32(total),
		},
	}
	return res
}
func Remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}
