package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	test1 "test/src/test"
)




// 同感`来给结构体的参数加上tag ，因为大写的话序列化出来的参数名是大写的，而小写的话无法被访问



type TestCar interface {
	Run()
	Stop()
}


type StrCar string

func (s StrCar) Run(){

}
func (s StrCar) Stop(){

}


type Benz struct {
	Name string
}

type BMW struct {
	Name string
}

func (benz Benz) Run() {
	println("benz run")
}

func (benz Benz) Stop() {
	println("benz stop")
}

func (benz Benz) GetName() string{
	println("benz name : " + benz.Name)
	return benz.Name
}

func (benz Benz) TestPrint() string{
	println("benz name : " + benz.Name)
	return benz.Name
}

func (bmw *BMW) Run() {
	println("bmw run")
}

func (bmw *BMW) Stop() {
	println("bmw stop")
}

func (bmw *BMW) GetName() string{
	println("bmw name : " + bmw.Name)
	return bmw.Name
}

func (bmw *BMW) Fly() {
	println("bmw fly")
}

type Person struct {

	Name string `json:"name"`
	Age uint8    `json:"age"`
}


func (p *Person) AddAge(num uint8) uint8 {

	p.Age += num
	return p.Age
}


type perss struct {
	test string
	testint int
}

func main()  {
	//du := Person{Name:"du" , Age:18}
	//toJsonFunc(du)
	//fmt.Println(du.Age)
	testInterface()
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




//‘ 结构体序列化
// 序列化别名问题 ： 加上`json:"别名"`
// 如果要修改结构体内部变量，传递指针 ，因为函数是值传递? 如果接受者类型是*结构体名字，那么传入的取地址，接受者的类型是地址，函数内外该地址对应的所有变量一起联动。
// 如果接受者类型是结构体名字，那么传入的和接受者的类型都是值，函数内外该变量值不一起联动。
// 指向指针的指针最终在序列化中能够被处理
//
func toJsonFunc(p Person)  {

	fmt.Println(p.Age)
	p.Age = 19
	fmt.Println(p.Age)
	//testp := Person{Name:"111"}
	//s ,_:= json.Marshal(&testp)
	//ss ,_:= json.Marshal(*p)
	//fmt.Println(string(s))
	//fmt.Println(string(ss))

	//var q **Person
	//q = &p
	//sss ,_:= json.Marshal(q)
	//fmt.Println(string(sss))
	////fmt.Println(*p)
	////fmt.Println(**q)
	//
	//var t ***Person
	//t = &q
	//ssss ,_:= json.Marshal(t)
	//fmt.Println("ssss" + string(ssss))
	//fmt.Println(*p)
	//fmt.Println(**q)

}


// 结构体方法
// 对结构体内部的值修改 / 不修改 如果方法使用的是指针则修改

func funcTest()  {

	p := Person{Name:"du" , Age:18}
	num :=p.AddAge(8)
	fmt.Println(num)
	fmt.Println(p)

}



//抽象




// 结构体的比较和匿名字段
// 如果内部变量完全一致则相等
// 如何判断两个变量是否指向同一个空间
func compareTest(){
	p := Person{Name:"du" , Age:18}
	q := Person{Name:"du" , Age:18}

	if p == q {
		println(true)
	}else {
		println(false)
	}

	//p.Age = 20

	if p == q {
		println(true)
	}else {
		println(false)
	}

	ptr1 := &p
	ptr2 := &q
	ptr1.Age = 20

	println(ptr1)
	println(ptr2)
	fmt.Println(&ptr1)
	fmt.Println(&ptr2)

}

// 多态

// 一个类必须实现接口的全部方法才算继承该接口 ， 否则编译报错
// 要实现多态的时候必须使用new的方法(除非实现接口的方法不采取指针的方式), 并且变量一开始申明为接口对应的接口类型 ， 应该是为了避免多个接口都有相同的方法，造成不知道是否实现的缘故
// 将对象复制给接口，会发生拷贝，而接口内部存储的是指向这个复制品的指针，即无法修改复制品的状态，也无法获取指针
//空接口可以作为任何类型数据的容器
// 接口在别的包如何调用
func testInterface() {

	var car test1.Car
	TestFuncParam(car)

	car = new(Benz)
	var itest interface{}


	car = Benz{Name:"1"}

	car.Run()
	car.Stop()
	fmt.Println(car.GetName())

	car = new(BMW)
	car.Run()
	car.Stop()
	fmt.Println(car.GetName())

	benz := Benz{Name:"111"}
	itest = car
	fmt.Println(itest)
	itest = benz
	fmt.Println(itest)

	TestFuncParam(Benz{Name:"111"})
	TestFuncParam(car)

	//TestFuncParam(Person{Name:"111"}) 如果没有实现该接口，编译无法通过
	//test := new(Car)
	//test = BMW{Name:"1"}
	//test

}

// 如果传进来的是未实现的接口实例 类型为nil ，如果是实现的，类型为对应的结构体
func TestFuncParam(car test1.Car) {
	fmt.Println("type:",reflect.TypeOf(car))
	fmt.Println(car)
}


// 内嵌 (结构体的内嵌类似继承)
// 结构体和接口的嵌入    同参数名如何解决 todo 内嵌基本数据类型会如何？
// 匿名结构体测试
// 通过类型 outer.int 的名字来获取存储在匿名字段中的数据，于是可以得出一个结论：在一个结构体中对于每一种数据类型只能有一个匿名字段。

func TestAnonymousStruct(){

	Config := struct{

		Code int
		Name string

	}{
		Code : 1 ,
		Name : "测试" ,
	}

	tempData := struct {
		HashTouchId bool
	}{
		HashTouchId: true,
	}

	fmt.Println(tempData.HashTouchId  )
	fmt.Println(Config)
}


// 序列化的时候如何把内嵌的序列化出来
// 反序列化的时候要传递的是指针  todo 如何辨别什么时候传递指针 Unmarshal传递的是interface{}
// 没办法通过内嵌对象实现多态 ， 但是可以配合上接口实现 https://www.cnblogs.com/li-mingxie/p/6928557.html
func TestAnonymousStructTwo(){

	apple := RedApple{}
	apple.Location = "bj"
	apple.Color = "red"
	apple.Shaper.Name = "no"
	//apple = new(GreenApple{})

	fmt.Println(apple)
	js , _ := json.Marshal(apple)
	fmt.Println(string(js))
	appleTwo := &RedApple{}
	s := ""
	json.Unmarshal(js , appleTwo)
	fmt.Println(appleTwo)
	json.Unmarshal(js , s)
	fmt.Println(s)
}

type Shape struct {
	Name string
}


type RedApple struct {
	Color string
	Apple
	Shaper Shape

}

type GreenApple struct {
	Color string
	Apple
}

type Apple struct {
	Location string
}

func (apple *Apple) getLocation (){
	println("apple location")
}

//func (apple *Apple) getLocation (i int) string{
//	println("apple location")
//}

func (redApple *RedApple) getLocation (){
	println("redApple location")
}

// 同一个包下不能有同名函数
//func testtest() {
//
//}
//
//func testtest(i int) string {
//
//}

// 方法可以重写 不能重载
func testRewrite() {

	apple := RedApple{}

	apple.getLocation()

}

// 方法是否可以重写和重载

// 当继承的时候有相同的方法

type A struct {

}

func (a *A)tt ()  {

	fmt.Println("a")

}

type B struct {

}

func (a *B)tt ()  {

	fmt.Println("b")

}

type C struct {
	A
	B
}

func testSameFunc()  {

	c := C{}
	c.B.tt()

}