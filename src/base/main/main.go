package main

import (
	"fmt"
	"os"
)

func main()  {

	testForLoop()

}



func printTest() {
	// go build和go run 失败 ：
	//原因在环境变量中的goroot配置的是用户目录下的go目录，但是这并不是go程序安装时的go执行文件的路径，程序重新将goroot配置为 /usr/lib/go即可

	// go install 失败 ： 原理同上 gopath的配置为 ～/go/src 但是文件是建立在 ～/go/下面  将gopath修改一下为 ～/go 更加合理

	fmt.Println("printTest")
}

// 带f的均可进行规格输出
// Fprintf 可用于输出流
func printCompare(){
	fmt.Println("Println")
	fmt.Printf("printf\n")
	fmt.Print("print")
	fmt.Printf("这是一个规格 %s , %s" , "test")
	sprintTest := fmt.Sprint("这是一个%s ， %s \n" , "测试" , "测试2\n")
	print(sprintTest)
	sprintTestOne := fmt.Sprintf("这是一个sprintTestOne, %s" ,"test\n" )
	print(sprintTestOne)
	fmt.Fprintf(os.Stdout , "this is %s\n" , "Fprint")
	fmt.Fprint(os.Stdout , "this is %s \n" , "Fprint")
	fmt.Fprint(os.Stderr , "this is os.stderr")
	//fmt.Sprint("111")
}


// 同：= 声明过的变量不能再声明 ， 如（ g := 10 g:=11），除非有新的参数用于声明（如 g := 10 , g,h := 10 , 10 (这种用法有问题)）

func testConstAndVar() {

	var a int
	var b string
	var c bool
	var d float32
	var e byte
	fmt.Printf("%d , %s , %t , %2.10f , %b \n" , a , b, c , d , e )

	var f = 10.1
	g := 20
	fmt.Printf("g in %d \n" , &g)
	g , h := 30 , 40
	fmt.Printf("g in %d \n" , &g)
	fmt.Println(f , g , h)

	const i , j , k = 10 , true , "sss"
	fmt.Println(i , j , k)

	const(
		l = iota
		m = iota
		n = iota
	)
	fmt.Println(l,m,n)


}

// var c = a + b 或者 c := a + b 在a ， b 类型不同（int8 int16）的时候报错 进行四则运算的两种类型必须相同

func testVarType()  {
	var a  int16
	var b int16
	a = 127
	b = 2
	b++
	c := a << b

	fmt.Printf("%T , %T , %T  \n" , a , b , c)
	fmt.Printf("%d , %d  %d %d \n" , a , b , c , a & b)

	var d float32
	var e float32
	d = 32
	e = 64
	f := d / e
	fmt.Printf("%T , %T , %T  \n" , d , e , f)
	fmt.Printf("%d , %d  %d %d \n" , d, e , f , d*e)


	var o complex64
	var p complex64
	o = 10 + 20i
	p = 30 + 50i
	q := o + p

	fmt.Printf("%T , %T , %T  \n" , o , p , q)
	fmt.Printf("%f , %f  %f %f \n" , o, p , q , o * p)

}




