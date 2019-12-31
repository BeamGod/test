package main

import (
	"fmt"
	"os"
)



func main()  {
	//
	m := 2
	k := funcOne(m)
	l := funcOne(4)
	fmt.Println("-------------------")
	fmt.Println(k)
	fmt.Println(l)
	fmt.Println("-------------------")
	fmt.Println(k())
	fmt.Println(k())
	fmt.Println(l())

	//testPointer()

	//a := []int{0 , 1 , 2 , 4 , 5}
	//b := a[1]
	//fmt.Println(a)
	//fmt.Println(b)

	//funcOne()

}



func printTest() {
	// go build和go run 失败 ：
	//原因在环境变量中的goroot配置的是用户目录下的go目录，但是这并不是go程序安装时的go执行文件的路径，程序重新将goroot配置为 /usr/lib/go即可

	// go install 失败 ： 原理同上 gopath的配置为 ～/go/src 但是文件是建立在 ～/go/下面  将gopath修改一下为 ～/go 更加合理

	fmt.Println("printTest")


}

// 带f的均可进行规格输出
// Fprintf 可用于输出流
// 格式化输入输出中的... 的作用 ： 函数有多个不确定参数的去情况下用于参数的传入  还可以作为切片的打散进行传递
func printCompare(){
	fmt.Println(len("Println"))
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
	c := a + b

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


func testFuncAsVar(a string) string{
	println(a)
	return a
}

// 闭包测试
// 闭包内的函数在打印出来的地址是一样的
// 函数的传入参数的是一个值copy ， 所以a的地址在同一次赋值，多次调用是一样的，同一个实例作为传入参数，不同次在外部赋值调用，地址不同
// 其中，如果闭包传给同一个参数 ， 这个参数多次调用 o的地址是一样的 ， 说明o在闭包内外是指向同一个空间
func funcOne(a int ) func() int{
	o := 3
	println(&o)

	return func() int {
		o++
		println(&o)
		println(&a)
		println(o + a)
		return o + a + o
	}
}


