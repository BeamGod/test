package main

import "fmt"

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
func loopSlice() {
	var str_slice []string

	//str_slice.

	println(str_slice)
	fmt.Println(str_slice)
}



// for循环两个list 排序d


// map slice 排序



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