package chat

import (
	"fmt"
	"net"
)

func RunDemo() {
	listen, err := net.Listen("tcp", "0.0.0.0:8080")

	if err != nil {
		panic(err.Error())
	}

	defer func(listen net.Listener) {
		err := listen.Close()
		if err != nil {
			panic(err.Error())
		}
	}(listen)

	for {
		fmt.Println("\n\n等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("发生 error: %v \n", err.Error())
		}

		fmt.Printf("有客户端连接上了 ip: %v\n", conn.RemoteAddr())

		msg := ""
		n, err := conn.Read([]byte(msg))
		if err != nil {
			fmt.Printf("conn Read 发生 error: %v \n", err.Error())
		}
		fmt.Printf("n == %v, readMsg === %v \n", n, msg)
	}
}
