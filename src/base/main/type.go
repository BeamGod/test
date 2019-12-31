package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strings"
)

func typeChange(){

	intToBytes()
}



// 直接通过[]byte将string强转获得的是[]byte，每个里面装的是单个字符的ascll码
// byte[]和string可以相互强制转换
// str 和[]byte之间可以直接相互转换

// str强转byte数组 测试 ok
//str转byte切片后判断容量 拓展容量  测试 ok
// byte切片转str后修改byte的值  ok 不影响string值 ，说明是将值copy过去
// 查看修改前后string的地址的变化  ok 地址没变化 说明string在修改前后指向的是相同的空间
// string 无法使用cap ok 说明string不是[]byte的别名
func byteSlice() {

	//var a byte
	//a = 'a'
	//fmt.Println(a)
	//b := int(a + 1)
	//fmt.Println(b)
	//c := string(b)
	//println("-------")
	//fmt.Println(c)
	//println("-------")
	//str2 := "hello"
	//fmt.Println(str2)
	//data2 := []byte(str2)
	//fmt.Println(data2)
	//str3 := string(data2)
	//fmt.Println(data2)
	//fmt.Println(str3)

	//var a []byte
	//fmt.Println(cap(a))
	//
	//a = append(a, '1')
	//fmt.Println(cap(a))
	//
	//a = append(a, '1')
	//fmt.Println(cap(a))
	//
	//a = append(a, '1')
	//
	//fmt.Println(a)
	//b := string(a)
	//fmt.Println(b)
	//
	//fmt.Println(cap(a))

	//s := "111"
	//fmt.Println(&s)
	//s = "111111111111111111111111111111111111111111111111111111111111111111111"
	//fmt.Println(&s)


	var a []byte
	fmt.Println(cap(a))

	a = append(a, '1')
	fmt.Println(cap(a))

	a = append(a, '1')
	fmt.Println(cap(a))

	a = append(a, '1')

	fmt.Println(a)
	b := string(a)
	fmt.Println(b)

	a = append(a, '1')
	a[0] = '0'
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(cap(a))




}

// 不能将byte数组转string,反之亦不行
func byteArr()  {
	//var c [3]byte
	//c[0] = 'a'
	//c[1] = 'b'
	//c[2] = 'c'
	//
	//s := string(c)
	//fmt.Println(s)
	//
	//s := "100"
	//a := [100]byte(s)

}

func testString(){
	str1:=[]string{"hello","world"}
	str:=strings.Join(str1,"")
	fmt.Println(str)
	fmt.Println(str[0:2])

}


// 不能直接将int强转为[]byte
// wirte函数中传入为 io类型 buff为×Buffer类型 为什么能够调用 似乎也没有继承相关接口 todo
func intToBytes() {
	var b []byte
	b = make([]byte , 0)
	a := 10
	buff := bytes.NewBuffer([]byte{})
	var d = 10
	binary.Write(buff, binary.BigEndian, int64(a))
	b = buff.Bytes()
	fmt.Println(a)
	fmt.Println(d)
	fmt.Println(b)

}

func bytesToInt(){

}

func intToString() {

}

func stringToInt() {

}

func stringToBytes(){

}

func bytesToString() {

}
