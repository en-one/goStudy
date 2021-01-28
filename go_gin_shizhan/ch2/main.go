package main

import "github.com/gin-gonic/gin"

type User struct {
	ID   uint64
	Name string
}

func main() {
	users := []User{{ID: 123, Name: "正经按摩的"}, {ID: 456, Name: "不长进干嘛"}}
	r := gin.Default()
	r.GET("/users", func(c *gin.Context) {
		c.JSON(200, users)
	})
	r.Run(":8090")
}
