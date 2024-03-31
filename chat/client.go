package chat

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func RunClientDemo() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Printf("connet error: %v", err.Error())
	}

	defer func(conn net.Conn) {
		conn.Close()
	}(conn)

	//读取终端输入
	reader := bufio.NewReader(os.Stdin)
	readString, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("os.Stdin error: %v", err.Error())
	}

	connByteCount, err := conn.Write([]byte(readString))
	if err != nil {
		fmt.Printf("conn write error: %v", err.Error())
	}

	fmt.Printf("客户端写入了 %v 字节的数据, 内容是: %v", connByteCount, readString)

}
