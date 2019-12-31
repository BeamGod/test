package main

import (
	"fmt"
	"sort"
	. "strconv"
)



var ttt int

func testForLoop()  {

	var i = 10

	for  j := i ; j >0 ; j-- {
		print(j)
	}

	for i >0 {
		print(i)
		i--
	}

	//Loop : for {
	//	print(i)
	//	i--
	//	if i> -100{
	//		goto Loop
	//	}
	//}
	
}


// for 循环字典 排序
//map类型中如果拿出不存在的值则发回改值的类型的零值
// 打印map带参数名
// print 和 fmt.println 打印出来的map 一个是地址 ， 一个是k v形式  println的有一个pointer的函数作为区分map chan等类型进行打印
//cap 不可以运用与获取map的长度 len可以 cap是获取容量的  len获取的是长度
func loopMap()  {
	//var name_age_map map[string]int
	//name_age_map = make(map[string]int)
	name_age_map := make(map[string]int)
	name_age_map["laow"] = 10
	name_age_map["a"] = 11
	name_age_map["b"] = 12
	name_age_map["c"] = 13
	print("\n")
	print(name_age_map["laow"])

	for key , value:= range name_age_map {
		fmt.Println(key)
		fmt.Println(value)
	}

	val , ok := name_age_map["c"]

	print("\n")
	print(val )
	print(ok)
	name_sex_map := make(map[string]bool)
	name_sex_map["a"] = true
	stringVal , ok2 := name_sex_map["a"]
	print("\n")
	print(stringVal )
	print(ok2)
	print("\n")

	name_class_map := map[string]string{"a" : "1" , "b" : "2" , "c" : "3"}
	println(name_class_map)
	delete(name_class_map , "p")
	fmt.Println(name_class_map)
	println("-----------")
	//intNum := 10
	rangeTest(name_age_map, name_class_map)

	fmt.Println(len(name_age_map))


}


// for 循环 排序数组
//cap和len都可以用于获取数组的长度 ，其实此处长度和容量相，因为数组是固定的
func loopArray()  {

	i := 0
	j := 100
	var arr [10]int


	for ; i < 5 ; i++ {
		arr[i] = i + j
		println(arr[i])
		fmt.Println(arr[i] + j)
	}


	for a , b := range  arr {
		println(a)
		println(b)

	}

	rangeTest(arr)
	fmt.Println(arr)


	var arr1 =  [...]float32{3.0 , 4.0}
	arr1[0] = 1.1
	fmt.Println(arr1)
	fmt.Println(len(arr))

}

//for 循环 slice  排序
// int 转string 直接用string（）会失败 ，需要使用Itoa函数 (string() 获取转的是对应的ascll码)
// 使用切片的时候，切出一部分给另一个切片，他们对应的值的空间是一样的（无论如何，切片本身的地址空间是不变的，即str1_slice 和str2_slice是不变额达），修改其中一个会影响另一个切片
//在使用append时，如果不超出其容量，则只和上面一样，但是如果使用append的时候，超出切片本身的cap，则切片本身指向地址不变，但是各个参数上指向的地址空间发生了变化，此时，修改一个切片不影响另一个切片

func loopSlice() {
	a := 97
	s := string(a)
	fmt.Println(s)
	var str_slice []string
	str1_slice := make([]string , 4 , 5)
	str1_slice[0] = "10"
	str1_slice[1] = "11"
	str1_slice[2] = "12"
	str1_slice[3] = "13"
	str2_slice := str1_slice[1:3]
	//println(&str1_slice)
	//println(&str2_slice)
	//println(&str1_slice)
	fmt.Println(str1_slice)
	fmt.Println(Itoa(len(str1_slice)))
	fmt.Println(Itoa(cap(str1_slice)))

	fmt.Println(str2_slice)
	fmt.Println(Itoa(len(str2_slice)))
	fmt.Println(Itoa(cap(str2_slice)))
	println(&str1_slice[1])
	println(&str2_slice[0])
	str2_slice = append(str2_slice, "5")
	println(&str1_slice[1])
	println(&str2_slice[0])
	str2_slice = append(str2_slice, "6")
	println(&str1_slice[1])
	println(&str2_slice[0])
	str2_slice = append(str2_slice, "7")
	println(&str1_slice[1])
	println(&str2_slice[0])
	//println(&str1_slice)
	//str1_slice = append(str1_slice, "7")
	println(&str1_slice)
	str2_slice[0] = "9"
	//str1_slice = append(str1_slice, "5")
	//str1_slice = append(str1_slice, "5")
	//str1_slice = append(str1_slice, "5")
	//str1_slice = append(str1_slice, "5")

	println(str_slice)
	fmt.Println(str_slice)

	println(str1_slice)
	fmt.Println(str1_slice)
	fmt.Println(Itoa(len(str1_slice)))
	fmt.Println(Itoa(cap(str1_slice)))

	fmt.Println(str2_slice)
	fmt.Println(Itoa(len(str2_slice)))
	fmt.Println(Itoa(cap(str2_slice)))
}



// for循环两个list 排序d


// map slice 排序

// map和slice在函数传递的时候是传引用还是值copy
//map 是传引用

func testSort() {

	map_test := make(map[string]string)
	map_test["3"] = "1"
	map_test["1"] = "1"
	map_test["2"] = "1"
	map_test["5"] = "1"
	map_test["0"] = "1"
	sortMap(map_test)

	fmt.Println(map_test)



	slice_tmp := make([]string , 4)
	slice_tmp[0] = "0"
	slice_tmp[1] = "1"

	fmt.Println("---------------")
	fmt.Println(&slice_tmp)
	fmt.Println(slice_tmp)
	fmt.Println(len(slice_tmp))
	fmt.Println(cap(slice_tmp))
	testSliceOne(slice_tmp)

	fmt.Println("---------------")
	fmt.Println(&slice_tmp)
	fmt.Println(slice_tmp)
	fmt.Println(len(slice_tmp))
	fmt.Println(cap(slice_tmp))


}




func sortMap(mp map[string]string) {

	slice_tmp := make([]string , 0)
	mp["0"] = "10"
	for key , _ := range mp {
		slice_tmp = append(slice_tmp, key)
	}
	fmt.Println(slice_tmp)
	sort.Strings(slice_tmp)
	fmt.Println(slice_tmp)

}


func testSliceOne(slice_tmp []string) {
	fmt.Println("---------------")

	fmt.Println(&slice_tmp)
	fmt.Println(slice_tmp)
	fmt.Println(len(slice_tmp))
	fmt.Println(cap(slice_tmp))
	slice_tmp[2] = "2"
	slice_tmp[3] = "3"
	slice_tmp = append(slice_tmp, "1")
	slice_tmp = append(slice_tmp, "2")
	slice_tmp = append(slice_tmp, "3")
	slice_tmp = append(slice_tmp, "4")
	fmt.Println("---------------")
	fmt.Println(&slice_tmp)
	fmt.Println(slice_tmp)
	fmt.Println(len(slice_tmp))
	fmt.Println(cap(slice_tmp))

}


// range 在数组 切片 int中的运用
// 此处range分别返回第n个数和其地址空间？
func rangeTest(a ...interface{})  {
	//var e *int
	//print(e)
	print("--------------------\n")
	for c , d := range  a {

		println(c)
		println(d)

	}
}




// 实现一个hashmap