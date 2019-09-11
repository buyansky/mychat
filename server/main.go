package main

import (
	"fmt"
	"github.com/buyansky/mychat/utils"
	"net"
	"sync"
)

var wp sync.WaitGroup

func main() {
	var key uint8
	fmt.Print("请输入密钥:")
	fmt.Scan(&key)

	listener, e := net.Listen("tcp", "0.0.0.0:8888")
	if e != nil {
		fmt.Printf("服务启动失败:%v\r\n",e)
		return
	}
	fmt.Printf("启动成功,地址:%v\r\n",listener.Addr())
	defer listener.Close()

	for {
		conn, e := listener.Accept()
		if e != nil {
			fmt.Printf("监听端口失败:%v\r\n",e)
			return
		}

		go utils.Read(conn,key)
		go utils.Write(conn,key)
	}

}
