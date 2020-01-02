package main

import (
	"fmt"
	"test/src/net/server"

	//"test/src/net/server"

	//"net"

	"log"
	//"net"
	"net/http"
	//"../server"
	//"../client"
)



func main()  {
	//netTest()
	server.StartServer()
	//client.StartClient()

}

func netTest()  {
	newServerMux()
	//net.Listen()
	//net.Dial()
	//net.Conn().
}


// tcp udp



func newServerMux()  {
	mux := http.NewServeMux()
	mux.HandleFunc("/", firstHandler)
	mux.HandleFunc("/hi", httpHiHandler)
	mux.HandleFunc("/hei/", httpHeiHandler)
	mux.HandleFunc("/hei/xxx", tmpHandler)
	err := http.ListenAndServe("localhost:8082" , mux)
	if err != nil {
		//fmt.Println()
		fmt.Println("err :")
		log.Println(err)
	}
}


// 同时进入hei和hi的handler

// 多级/是否会出问题
// 设置默认handler
func listenServer(){
	http.HandleFunc("/", firstHandler)
	http.HandleFunc("/hi", httpHiHandler)
	http.HandleFunc("/hei/", httpHeiHandler)
	http.HandleFunc("/hei/xxx", tmpHandler)
	err := http.ListenAndServe("localhost:8081" , nil)
	if err != nil {
		//fmt.Println()
		fmt.Println("err :")
		log.Println(err)
	}
	//net.Listen()
}

func firstHandler(response http.ResponseWriter, request *http.Request)  {
	fmt.Println("进入handler方法了 first")
	url := request.URL.Path
	fmt.Println(url)
	//methodType := request.Method
	//log.Println(methodType)
	hosts := request.URL.Host
	log.Println(hosts)
	response.Write([]byte("22222"))
	response.Header()
}


func tmpHandler(response http.ResponseWriter, request *http.Request)  {
	fmt.Println("进入handler方法了 tmpHandler")
	url := request.URL.Path
	fmt.Println(url)
	//methodType := request.Method
	//log.Println(methodType)
	hosts := request.URL.Host
	log.Println(hosts)
}

func httpHiHandler(response http.ResponseWriter, request *http.Request){
	fmt.Println("进入handler方法了 hi")
	url := request.URL.Path
	fmt.Println(url)
	//methodType := request.Method
	//log.Println(methodType)
	hosts := request.URL.Host
	log.Println(hosts)
}

func httpHeiHandler(response http.ResponseWriter, request *http.Request)  {
	fmt.Println("进入handler方法了 hei")
	url := request.URL.Path
	fmt.Println(url)
	//methodType := request.Method
	//log.Println(methodType)
	hosts := request.URL.Host
	log.Println(hosts)
}