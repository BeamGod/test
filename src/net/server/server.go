package server

import (
	"net"
	"log"
)

// 保持消息
var MessageQueue = make(chan string , 1024)
//var ip_conn_map = make(map[string]net.Conn)
var conn_slice = make([]net.Conn , 0)

// 记录用户ip和对应的链接
func StartServer()  {


	socketListen , err := net.Listen("tcp" , "127.0.0.1:8081")
	if err != nil {
		log.Println("链接发生错误")
	}
	defer socketListen.Close()

	for {
		conn , err := socketListen.Accept()
		conn_slice = append(conn_slice, conn)
		if err != nil {
			log.Println("accpet 错误")
		}
		//ip_conn_map[conn.RemoteAddr().String()]  = conn
		defer conn.Close()
		// 将ip记录到对应链接

		go consumeMsg()
		go accpetMsg(conn)

	}

	
}


func accpetMsg(conn net.Conn)  {

	ip := conn.RemoteAddr().String()
	log.Println("ip :" + ip)
	buf := make([]byte , 1024)
	for {
		num, err := conn.Read(buf)
		if err != nil {
			//log.Println("accpetMsg 错误" , err)
		}
		MessageQueue <- string(buf[0:num])
	}

}

func consumeMsg()  {

	for {
		msg ,ok:= <- MessageQueue
		if !ok {
			continue
		}

		for _ , conn := range conn_slice{
			log.Println("accpetMsg 错误" , conn)
			conn.Write([]byte(msg))

		}
	}

}
