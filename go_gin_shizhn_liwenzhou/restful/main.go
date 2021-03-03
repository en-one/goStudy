package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// /book查询书籍信息
	r.GET("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Get",
		})
	})

	// /post 创建书籍记录
	r.POST("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "POST",
		})
	})

	r.PUT("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "PUT",
		})
	})

	r.DELETE("/book", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "DELETE",
		})
	})

	r.Run(":8090")

}
