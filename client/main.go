package main

import (
	"fmt"
	"github.com/buyansky/mychat/utils"
	"net"
	"sync"
)
var wp sync.WaitGroup


func main() {
	var Addr string
	var Key uint8
	fmt.Print("请输入服务端ip和端口:")
	fmt.Scan(&Addr)
	fmt.Print("请输入密钥:")
	fmt.Scan(&Key)



	conn, e := net.Dial("tcp", Addr)
	if e != nil {
		fmt.Printf("连接服务端失败:%v\r\n",e)
		return
	}

	wp.Add(2)
	go utils.Read(conn,Key)
	go utils.Write(conn,Key)
	wp.Wait()

}



