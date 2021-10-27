package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/user/:uid", handleUser)
	r.Run()
}

func handleUser(c *gin.Context) {
	uid := c.Param("uid")
	// TODO: dynamoからid = uidのユーザ名を取得する
	// TODO: テーブル名を環境変数からとる -> CDK側で実装する必要あり
	tableName := os.Getenv("DYNAMODB_TABLE")
	descdynamodb(&tableName)
	c.JSON(200, gin.H{
		"message": fmt.Sprintf("Hello %s", uid),
	})
}
