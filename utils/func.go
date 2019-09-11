package utils

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func Read(conn net.Conn,key uint8)  {
	defer conn.Close()
	for{
		reader := bufio.NewReader(conn)

		var buf [128]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Printf("失败:%v\r\n",err)
			break
		}
		data:=Decode(buf[:n],key)
		fmt.Printf("收到信息: %s\r\n",data)

	}

}


func Write(conn net.Conn,key uint8){
	defer conn.Close()
	newReader := bufio.NewReader(os.Stdin)

	for{
		s, _ := newReader.ReadString('\n')
		trim := strings.Trim(s, "\r\n")
		info:=Encode([]byte(trim),key)
		conn.Write(info)
	}


}