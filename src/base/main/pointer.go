package main

import "fmt"

func testPointer() {

	var (
		name string
		age int
	)

	fmt.Scanf("1:%s 2:%d ", &name, &age)
	fmt.Println(age , name)

}

func testPointerOne() {

	a := 10
	p := &a
	var q *int

	fmt.Println(q)

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Printf("%x\n" ,&a)
	fmt.Printf("%x\n" ,p)

	fmt.Println(*p)
	fmt.Println(&p)
}

// 指针数组

//数组指针

// 函数指针传参


// go中没办法简单实现指针的运算 ， 需要通过uintptr  unsafe.Pointer 等方法去操作
func testPointerArr() {

	a := []int{ 1 , 2 ,3}
	fmt.Println(a)
	var p [3]*int

	for i , j := range a {
		fmt.Println(i)
		fmt.Println(j)
		p[i] = &a[i]
	}


	//for i := range p {
	//	print(*p[i])
	//}

	//var q *int
	//q = &a[0]
	//
	//for i := 0 ; i < 3 ; i++ {
	//	fmt.Println(q)
	//	q++ 1
	//}
}


func testArrPointer() {
	a := []int{ 1 , 2 ,3}

	var p *[]int

	p = &a

	println(a)
	println(&a)
	println(p)
	println(&p)
	println(*p)
	fmt.Println("--------------")
	println((*p)[0])
	println(a[0])

	fmt.Println("--------------")

	fmt.Println(a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)
	fmt.Println(&a)
	fmt.Println("--------------")
	fmt.Println((*p)[0])

}

func pointFun(p *int)  {


	fmt.Println(p)

	fmt.Println(p)

	fmt.Println(*p)



}






