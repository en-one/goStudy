package main

import "github.com/gin-gonic/gin"

/*
分组路由
*/

func main() {
	r := gin.Default()

	//v1版本api
	v1Group := r.Group("/v1")
	{
		v1Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/users")
		})

		v1Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v1/products")
		})
	}
	//v2版本api
	v2Group := r.Group("/v1")
	{
		v2Group.GET("/users", func(c *gin.Context) {
			c.String(200, "/v1/users")
		})

		v2Group.GET("/products", func(c *gin.Context) {
			c.String(200, "/v1/products")
		})
	}

	r.Run(":8090")
}
