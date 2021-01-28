package main

/*
	中间件:
		强调的是中间，是一种无业务的，在正常的业务handler处理前后的 独立的逻辑片段
	调用顺序：serverMux路由分发->调用中间件1->调用中间件2......->调用真正的业务处理逻辑
	中间件非常独立，不用的时候 去掉即可


*/

//import  "github.com/gorilla/handlers"
import _ "github.com/gorilla/handlers"

//Gorilla Handlers 使用
func main() {

}
