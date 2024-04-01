package chat

import (
	"fmt"
	log "go-demo/logger"
	"net"
	"regexp"
	"strings"
)

// EXIT_FLAG 退出指令
const EXIT_FLAG = ":exit"

// CONNECT_IP_OPT 连接其他ip操作指令，例如：->192.168.5:8080:{这是消息}
const CONNECT_IP_OPT = "->"

// MAX_CONN_COUNT 最大连接 20 个
const MAX_CONN_COUNT = 20

type IPAddrMap = map[string]net.Conn

func progress(conn net.Conn, ipMap *IPAddrMap) {

	connIp := conn.RemoteAddr().String()

	defer func() {
		fmt.Printf("%v 连接已断开....\n", connIp)
		// 连接退出时，从ip映射表中删除对应ip
		delete(*ipMap, connIp)
		conn.Close()
	}()

	for {
		fmt.Printf("[info] 等待客户端 %v 输入\n", connIp)

		//读取客户端输入的内容
		msgBuf := make([]byte, 1024)
		readByteCount, err := conn.Read(msgBuf)
		if err != nil {
			fmt.Printf("[error] 客户端连接发生异常：%v \n", err.Error())
			return
		}
		//将内容从字节转成string
		msg := string(msgBuf[:readByteCount])
		trimmedMsg := strings.Trim(msg, " \r\n")
		//收到退出指令
		if trimmedMsg == EXIT_FLAG {
			return
		}

		isConnIp := trimmedMsg[0:2] == CONNECT_IP_OPT
		// 连接其他ip
		if isConnIp {
			//正则匹配 ip 端口号 消息
			re := regexp.MustCompile(`->(\d+\.\d+\.\d+\.\d+):(\d+):(.*)`)
			//匹配结果
			match := re.FindStringSubmatch(trimmedMsg)
			//连接的ip+端口号
			connIp := match[1] + ":" + match[2]
			//要发送的消息
			sendMsg := match[3]
			// 从 ip 表中取connect实例
			optConn := (*ipMap)[connIp]
			//给对应的ip发送消息
			_, err := optConn.Write([]byte(sendMsg))
			if err != nil {
				log.Logger.Warn("发送消息失败", err.Error())
				return
			}
		}

		// 对消息做处理
		fmt.Printf("[reader] conn 读取到了消息：%v", msg)
	}
}

func RunDemo() {
	listen, err := net.Listen("tcp", "0.0.0.0:8080")

	ipAndConnMap := make(IPAddrMap, MAX_CONN_COUNT)

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

		remoteAddrStr := conn.RemoteAddr().String()
		fmt.Printf("有客户端连接上了，ip是: %v\n", remoteAddrStr)

		//保存用户的 ip 和 connect 实例
		ipAndConnMap[remoteAddrStr] = conn

		//开一个协程处理用户连接
		go progress(conn, &ipAndConnMap)
	}
}
