package chat

import (
	"fmt"
	"net"
)

// EXIT_FLAG 退出指令，注意客户端发送的时候，需要按回车 最后需要加上 \n
const EXIT_FLAG = ":exit\n"

func progress(conn net.Conn) {
	defer func() {
		fmt.Printf("%v 连接已断开....", conn.RemoteAddr().String())
		conn.Close()
	}()

	for {
		fmt.Printf("[info] 等待客户端 %v 输入\n", conn.RemoteAddr().String())

		//读取客户端输入的内容
		msgBuf := make([]byte, 1024)
		readByteCount, err := conn.Read(msgBuf)
		if err != nil {
			fmt.Printf("[error] 客户端连接发生异常：%v \n", err.Error())
			return
		}
		//将内容从字节转成string
		msg := string(msgBuf[:readByteCount])

		//收到退出指令
		if msg == EXIT_FLAG {
			return
		}

		// 对消息做处理
		fmt.Printf("[reader] conn 读取到了消息：%v", msg)
	}
}

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
		fmt.Println("\n等待客户端连接...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept 发生 error: %v \n", err.Error())
		}

		fmt.Printf("有客户端连接上了，ip是: %v\n", conn.RemoteAddr())

		//开一个协程处理用户连接
		go progress(conn)
	}
}
