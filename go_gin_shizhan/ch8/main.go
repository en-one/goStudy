package main

import "github.com/gin-gonic/gin"

type user struct {
	ID   int
	Name string
	Age  int
}

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		//c.JSON(200, gin.H{"message": "hellow world"})
		//c.JSON(200, user{ID: 123, Name: "张三", Age: 20})
		c.IndentedJSON(200, user{ID: 456, Name: "李四", Age: 25})
	})

	r.Run(":8090")
}

/*
输出json数组
*/
