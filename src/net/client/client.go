package client

import (
	"fmt"
	"log"
	"net"
)

var msgQueue =  make(chan string , 1024)

func StartClient()  {

	//socketListen , err := net.Listen("tcp" , "127.0.0.1:8082")
	socketListen , err := net.Dial("tcp" , "127.0.0.1:8081")
	if err != nil {
		log.Println(err)
	}
	defer socketListen.Close()
	log.Println("客户端和服务端链接")
	go sendMsg(socketListen)
	go getMsg(socketListen)

	for {
		s := ""
		fmt.Scanln(&s)
		if s == "" {
			continue
		}
		msgQueue <- s

	}




}


// 发送消息

func sendMsg(conn net.Conn)  {
	for  {
		buf := <- msgQueue
		conn.Write([]byte(buf))
	}
}

// 接收消息

func getMsg(conn net.Conn) {
	for   {
		buf := make([]byte , 1024)
		_ , err := conn.Read(buf)
		if err != nil {
			continue
		}
		fmt.Println("接收到消息 ： \n" , string(buf))
	}
}
