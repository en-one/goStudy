package main

import (
	"josiah.top/go_lagou/ch22/server"
	"log"
	"net"
	"net/rpc"
)

/*
rpc 远程服务 c/s

客户端（Client）调用客户端存根（Client Stub），同时把参数传给客户端存根；
客户端存根将参数打包编码，并通过系统调用发送到服务端；
客户端本地系统发送信息到服务器；

服务器系统将信息发送到服务端存根（Server Stub）；
服务端存根解析信息，也就是解码；
服务端存根调用真正的服务端程序（Sever）；

服务端（Server）处理后，通过同样的方式，把结果再返回给客户端（Client）。
*/

func main() {
	rpc.RegisterName("MathService", new(server.MathService))
	//rpc.HandleHTTP()//新增变为通过http协议调用
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}
	rpc.Accept(l)
	//http.Serve(l, nil)//换成http服务
}
