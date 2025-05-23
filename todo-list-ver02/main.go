package main

import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/khaquachtrong74/simple-api/net-http/middleware"
)
type myForm struct{
	Task string `form:"task"`
}

func formHandler(c *gin.Context){
	var fakeForm myForm
	fakeForm.Task = c.PostForm("task")
}
func main(){
//Custom HTTP Configuration
	router := gin.Default()
	router.LoadHTMLGlob("./*")
	router.GET("/task", func(c *gin.Context){
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "test",
		})
	})
	router.POST("/list", func (c *gin.Context){
		formHandler(c)
	})
	s := http.Server{
		Addr: ":8080",
		Handler: router,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}
