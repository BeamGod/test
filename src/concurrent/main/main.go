package main

import (
	"fmt"
	"sync"
	"time"
)

func main()  {
	//fmt.Prin
	//tln("111")
	testSelect()

}



// 协程中的i没打完 主线程先结束了 goroutine 最大申请1G栈
func goTest() {

	a := func(){

		go func() {
			for true{
				fmt.Println("test")
			}

		}()
		for i := 0 ; i < 1000 ; i++ {
			fmt.Println(i)

		}
		fmt.Println("over2")

	}


	go a()
	a()
	fmt.Println("over1")
}


// sync
type User struct {
	sync.Mutex
	sync.RWMutex

	name string
}


func testSyb(){
	//u1 := &User{name: "test"}
	//u1.Mutex.Lock()
	//defer u1.Mutex.Unlock()
	//
	//tmp := *u1
	//u2 := &tmp
	//u2.Mutex = sync.Mutex{} // 没有这一行就会死锁
	//
	//fmt.Printf("%v\n", u1)
	//fmt.Printf("%v\n", u2)
	//
	//u2.Mutex.Lock()
	//defer u2.Mutex.Unlock()

	u1 := &User{name: "test"}
	u1.RWMutex.RLock()
	defer u1.RWMutex.RUnlock()

	tmp := *u1
	u2 := &tmp
	//u2.RWMutex = sync.RWMutex{} // 没有这一行就会死锁

	fmt.Printf("%v\n", u1)
	fmt.Printf("%v\n", u2)

	u2.RWMutex.RLock()
	defer u2.RWMutex.RUnlock()
}


var wg sync.WaitGroup

// 当add的数量多过实际done的数量，会引起程序恐慌
func testGroupSyn(){



	wg.Add(2)
	go printA()
	go printB()

	fmt.Println("开始阻塞")
	wg.Wait()
	fmt.Println("结束阻塞")

}

func printA(){
	defer 	wg.Done()

	for i := 0 ; i< 100 ; i++ {
		fmt.Println("A : " , i)
	}

}

func printB(){
	defer 	wg.Done()

	for i := 0 ; i< 100 ; i++ {
		fmt.Println("A : " , i)
	}

}

// channel
// chan 读写都会阻塞 ，在size数量之前
func testChannel(){

	a := make(chan int , 2)
	var s string
	go func() {
		for i := 0 ; i < 50 ; i++ {
			a <- 1
			fmt.Println("in : " , i)
		}
	}()

	go func() {
		for i := 0 ; i < 50 ; i++ {
			<- a
			fmt.Println("out : " , i)
		}
	}()

	time.Sleep(time.Second * 5)
	fmt.Println(s)

}


// 当读取数据写入的数据结束的时候要调用close方法 ， 否则读取数据会报错(应该是程序处理，这里应该会阻塞才对)

func testChannelTwo() {

	ch := make(chan int , 1)

	go func() {
		for i := 0 ; i < 5 ; i++ {
			ch <- i
			//if i == 4 {
			//	close(ch)
			//}
			//if i == 5
			//printB()
		}
		close(ch)
	}()
	//
	//for true {
	//	v , ok := <- ch
	//	if !ok {
	//		fmt.Println("数据读取结束")
	//		break
	//	}else {
	//		fmt.Println(v)
	//	}
	//	time.Sleep(time.Second * 1)
	//}
	time.Sleep(time.Second * 1)

	for v := range ch {

		fmt.Println(v)
	}

	fmt.Println("1111")

}


// select todo
// chan 在关闭的时候会发送一个默认值，sekect是可以接受到的
func testSelect()  {

	ch1 := make(chan int )
	ch2 := make(chan int )

	go func() {
		time.Sleep(time.Second * 1)

		ch1 <- 100
	}()

	go func() {
		//time.Sleep(time.Second * 1)
		//close(ch2)
		//ch2 <- 200
	}()
	select {
	case num := <- ch1:
		fmt.Println(num)
	case num , ok:= <- ch2:
		fmt.Println(num , ok)

	}

}


func testChannelThr() {
	ch := make(chan int)
	go readOnly(ch)

	go writeOnly(ch)
	time.Sleep(time.Second * 1)

}

func readOnly( ch <- chan int)  {

	data := <- ch
	fmt.Println(data)

}

func writeOnly(ch chan <- int )  {
	ch <- 1

}