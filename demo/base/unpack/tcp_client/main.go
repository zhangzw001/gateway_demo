package main

import (
	"fmt"
	"github.com/zhangzw001/gateway_demo/demo/base/unpack/unpack"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9090")
	defer conn.Close()
	if err != nil {
		fmt.Printf("connect failed, err : %v\n", err.Error())
		return
	}
	unpack.Encode(conn, "hello world 0!!!")
}
