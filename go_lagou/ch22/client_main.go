package main

import (
	"fmt"
	"josiah.top/go_lagou/ch22/server"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234") //建立tcp链接，ip和端口要和rpc提供的服务一致
	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := server.Args{A: 7, B: 8}

	var reply int
	err = client.Call("MathService.Add", args, &reply)
	if err != nil {
		log.Fatal("MathService.Add error :", err)
	}
	fmt.Printf("MathService.Add: %d+%d=%d", args.A, args.B, reply)
}
