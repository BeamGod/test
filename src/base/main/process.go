package main

import (
	"fmt"
	"reflect"
)

// 同一个包内不同函数的相互调用
// 问题 ： 在同一个包内，main函数调用此函数，用编译器run可以， 但是用命令build， run有问题
//解决 ： 虽然用编译器可以帮我编译目录下全部文件，但是用命令的时候是不知道哪些文件是需要编译的 。 所以需要 go build main.go process.go  或者直接编译运行该package go run test/src/base/main 或者直接go run *.go

func testProcessControlller()  {

	if a , b := testFun() ; b{
		//printTest()
		print("a")
	} else {
		printTest()
		print(a)

	}

}

func testFun() (bool , bool) {
	return true , false
}


/*
问题 如果x为x interface{}时是可以获取变量类型的 x.(type)，但是为int的时候无法获取 Invalid type switch guard: i := x.(type) (non-interface type int on left)
解决 ： i := interface{}(x).(type) 强转类型  其中.(type) 是在switch中才有的，否则会报错Use of .(type) outside type switch
拓展 ： go中的类型断言 l , ok := interface{}(x).(string)
https://blog.csdn.net/wuhuimin521/article/details/83066417 类型断言

https://blog.csdn.net/u011957758/article/details/81193806 反射


 */


func testSwitch()  {

	var x interface{}
	x = "111"
	l , ok := interface{}(x).(string)
	//l , ok := interface{}x.(string)
	print(l)
	print(ok)


	print("\n")

	switch i := (x).(type) {   //switch i := x  报错 不理解（） todo
	case nil:
		fmt.Printf(" x 的类型 :%T",i)
	case int:
		fmt.Printf("x 是 int 型")
	case float64:
		fmt.Printf("x 是 float64 型")
	case func(int) float64:
		fmt.Printf("x 是 func(int) 型")
	case bool, string:
		fmt.Printf("x 是 bool 或 string 型" )
	default:
		fmt.Printf("未知型")
		printTest()
	}


}


func typeof(v interface{}) interface{} {
	return reflect.TypeOf(v)
}