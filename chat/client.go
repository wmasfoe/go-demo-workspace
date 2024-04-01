package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const EXIT_FLAG = ":exit"

// 开一个 goroutine 用来读取服务端的消息
func readFromServer(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Printf("conn read error: %v \n", err.Error())
			return
		}

		fmt.Printf("有其他用户给你发消息了: %v \n", string(buf[:n]))
	}
}

func RunClientDemo() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("连接服务端失败")
		panic(err.Error())
	}

	defer func(conn net.Conn) {
		fmt.Printf("%v 关闭客户端连接", conn.LocalAddr().String())
		conn.Close()
	}(conn)

	// 用来处理其他 client 发送的消息
	go readFromServer(conn)

	for {
		fmt.Printf("请输入要发送的内容(输入 %v 退出聊天): \n", EXIT_FLAG)
		//读取终端输入
		reader := bufio.NewReader(os.Stdin)
		readString, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("os.Stdin error: %v", err.Error())
		}

		connByteCount, err := conn.Write([]byte(readString))
		if err != nil {
			fmt.Printf("conn write error: %v", err.Error())
			return
		}

		fmt.Printf("客户端写入了 %v 字节的数据, 内容是: %v", connByteCount, readString)

		if strings.Trim(readString, " \r\n") == EXIT_FLAG {
			return
		}
	}

}
