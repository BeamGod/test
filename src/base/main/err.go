package main

import (
	"errors"
	"fmt"
)

type MyErr struct {
	err string
	code int
}

func (error *MyErr)Error() string {
	error.err = "错误类型为 %d"
	return fmt.Sprintf(error.err , error.code)

}

func testError ()  {
	panicAndRecoverTest()

}

func cucl()(v int , err error){
	if v == 0 {
		err = errors.New("函数错误")
		return
	}else{
		v =1
		return
	}
}


// defer 执行 类似后进先出
func deferIt3() {
	f := func(i int) int {
		fmt.Printf("%d ",i)
		return i * 10
	}
	for i := 1; i < 5; i++ {
		defer fmt.Printf("%d ", f(i))
	}
}


func deterIt1() {
	deferIt2()

	defer fmt.Println("1")
}

func deferIt2()  {
	defer fmt.Println("2")

}

func returnValues() int {
	var result int
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}

func namedReturnValues() (result int) {
	defer func() {
		result++
		fmt.Println("defer")
	}()
	return result
}


//panic recover

func panicAndRecoverTest()  {
	deferPanic()
}

func deferPanic() {

	defer fmt.Println("defer 1")
	panicFun()
	fmt.Println("run ")
	defer fmt.Println("defer 4")

}

func panicFun(){
	defer func() {
		if msg := recover() ; msg != nil {
			fmt.Println("recover----")
		}
	}()
	defer fmt.Println("defer 2")
	//panic("panic -----")
	for i := 3 ; i >= 0 ; i-- {
		if i == 0 {
			break
		}
		println(1 / i)

	}
	defer fmt.Println("defer 3")
}

