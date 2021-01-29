package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
restful api
---api规范
		GET，表示读取服务器上的资源；
		POST，表示在服务器上创建资源；
		PUT，表示更新或者替换服务器上的资源；
		DELETE，表示删除服务器上的资源；
		PATCH，表示更新 / 修改资源的一部分。
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

	http.HandleFunc("/users", handleUsers)
	http.ListenAndServe(":8090", nil)

}

func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": //只有get才能获取数据，才算是restful
		users, err := json.Marshal(users)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "{\"message\": \""+err.Error()+"\"}")
		} else {
			w.WriteHeader(http.StatusOK)
			w.Write(users)
		}
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "{\"message\": \"not found\"}")

	}

}
