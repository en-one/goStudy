package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "111"},
	{ID: 2, Name: "222"},
	{ID: 3, Name: "333"},
}

/*
gin 删除以及patch 部分更新
*/
func main() {
	r := gin.Default()
	r.DELETE("/users/:id", deleteUser)
	r.Run(":8090")
}

//-----------删除用户
func deleteUser(c *gin.Context) {
	id := c.Param("id")
	i := -1
	for index, u := range users {
		if strings.EqualFold(id, strconv.Itoa(u.ID)) {
			i = index
			break
		}
		if i >= 0 {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusNoContent, "")
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "用户不存在",
			})
		}
	}

}
