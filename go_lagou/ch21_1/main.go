package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
gin实现restful
*/

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

var users = []User{
	{ID: 1, Name: "111"},
	{ID: 2, Name: "222"},
	{ID: 3, Name: "333"},
}

func main() {
	r := gin.Default()
	//r.GET("/users",listUser)
	//r.GET("/users/:id", getUser)
	r.POST("/users", createUsesr)
	r.Run(":8090")
}

//--------------获取用户GET
//func listUser(c *gin.Context)  {
//	c.JSON(200, users)
//}
//-----------------获取特定的用户
//func getUser(c *gin.Context)  {
//	id:= c.Param("id")
//	var user User
//	found := false
//	for _,u := range users {
//		if strings.EqualFold(id, strconv.Itoa(u.ID)){
//			user = u
//			found = true
//			break
//		}
//	}
//	if found {
//		c.JSON(200, user)
//
//	}else {
//		c.JSON(404, gin.H{
//			"message":"用户不存在",
//		})
//	}
//}
//---------------------创建用户
func createUsesr(c *gin.Context) {
	name := c.DefaultPostForm("name", "")
	if name != "" {
		u := User{ID: len(users) + 1, Name: name}
		users = append(users, u)
		c.JSON(http.StatusCreated, u)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"mesage": "请输入用户名",
		})
	}

}
