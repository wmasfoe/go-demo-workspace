package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

const CLIENT_EXIT_FLAG = ":exit"

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

	for {
		fmt.Println("请输入要发送的内容(输入 :exit 退出聊天): ")
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

		if strings.Trim(readString, " \r\n") == CLIENT_EXIT_FLAG {
			return
		}
	}

}
