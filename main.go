package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Hello(ctx *gin.Context) {
	resp := &Response{
		Code: 0,
		Msg:  "succ",
		Data: fmt.Sprintf("hello world! %s,", time.Now()),
	}
	ctx.JSON(http.StatusOK, resp)
}
func main() {
	router := gin.Default()
	router.GET("/", Hello)
	err := router.Run(":8888")
	if err != nil {
		fmt.Println("ddd")
	}
}
