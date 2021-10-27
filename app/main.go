package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:uid", func(c *gin.Context) {
		uid := c.Param("uid")
		// TODO: dynamoからid = uidのユーザ名を取得する
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Hello %s", uid),
		})
	})
	r.Run()
}
