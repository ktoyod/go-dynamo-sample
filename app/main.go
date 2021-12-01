package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

var tableName = os.Getenv("DYNAMODB_TABLE")

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/description", hanldeGetDescription)
	r.GET("/user", handleGetUsers)
	r.GET("/user/:uid", handleGetUser)
	r.POST("/user", handlePostUser)
	r.Run()
}

func hanldeGetDescription(c *gin.Context) {
	descDynamoDB(&tableName)
	c.JSON(200, gin.H{
		"message": "Check AWS CloudWatch Logs",
	})
}

func handleGetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ユーザ一覧が返る予定",
	})
}

func handleGetUser(c *gin.Context) {
	uid := c.Param("uid")
	// TODO: dynamoからid = uidのユーザ名を取得する
	items := scanDynamoDB(&tableName, &uid)
	c.JSON(200, gin.H{
		"items": items,
	})
}

func handlePostUser(c *gin.Context) {
	// TODO: 新規作成するユーザのIDを決定
	// TODO: ペイロードからユーザ名を取得
	// TODO: ユーザ登録
	c.JSON(200, gin.H{
		"message": "ユーザ登録できたよ",
	})
}
