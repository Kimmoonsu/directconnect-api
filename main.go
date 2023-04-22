package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/", func(c *gin.Context) {
		// Request Header 출력
		fmt.Println("=============HEADER==========")
		fmt.Println(c.Request.Header)
		fmt.Println("")

		// Request Body 출력
		body, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			// 에러 처리
		}
		fmt.Println("==============BODY==============")
		fmt.Println(string(body))
		fmt.Println("")

		// Response
		c.String(http.StatusOK, "Success")
	})

	r.Run(":5000")
}
