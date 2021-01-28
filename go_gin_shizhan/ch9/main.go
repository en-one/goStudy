package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.SecureJSON(200, gin.H{"wechat": "josiah"})
	})

	a := []string{"1", "2", "3"}
	r.GET("/secureJson", func(c *gin.Context) {
		c.SecureJSON(200, a)
	})

	r.Run(":8090")
}

/*
跨域
使用jsonp
*/
