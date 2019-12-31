package main

import (
	"encoding/json"
	"fmt"
)


// 同感`来给结构体的参数加上tag ，因为大写的话序列化出来的参数名是大写的，而小写的话无法被访问
type Person struct {

	Name string `json:“name”`
	Age uint8    `json:”age“`
}


type perss struct {
	test string
	testint int
}

func main()  {
	du := Person{Name:"du" , Age:18}
	toJsonFunc(&du)
	//fmt.Println(du.Age)
}


// 结构体中参数首字母没有大写时，别的包虽不可以调用
func printStruct(){
	//fmt.Println(Person{name : "kkk" , age : 8})
	du := Person{Name:"du" , Age:18}
	//fmt.Println(du)
	fmt.Printf("%+v\n" , du)

	duperss := perss{test : "1" , testint : 1}
	fmt.Printf("%+v\n" , duperss)


	var p *Person
	p = &du
	fmt.Println(p.Age)

	var m Person
	var n *Person=  new(Person)
	var l *Person = &Person{}

	fmt.Println(m , n.Age, l.Age)


}


// 结构体方法


//‘ 结构体序列化
// 序列化别名问题 ： todo
// 如果要修改结构体内部变量，不许传递指针 ，因为函数是值传递?todo
// 指向指针的指针最终在序列化中能够被处理
func toJsonFunc(p *Person)  {

	fmt.Println(p.Age)
	p.Age = 19
	fmt.Println(p.Age)
	//testp := Person{Name:"111"}
	//s ,_:= json.Marshal(&testp)
	//ss ,_:= json.Marshal(*p)
	//fmt.Println(string(s))
	//fmt.Println(string(ss))

	var q **Person
	q = &p
	sss ,_:= json.Marshal(q)
	fmt.Println(string(sss))
	//fmt.Println(*p)
	//fmt.Println(**q)

	var t ***Person
	t = &q
	ssss ,_:= json.Marshal(t)
	fmt.Println("ssss" + string(ssss))
	//fmt.Println(*p)
	//fmt.Println(**q)

}

// 多态

//抽象

// 内嵌

// 结构体的比较和匿名字段