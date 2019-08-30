/*
   字典定义，操作，以及应用演示
*/

package main

import "fmt"

func main()  {
	dict04()
}

//字典的创建，标准格式
func dict01()  {
	var dict map [string] string= map[string] string{}
	dict["age"]="10"
	fmt.Print(dict["age"])
}

//字典创建，标准格式，没有初始化
func dict02()  {
	var dict2 map[string] string
	dict2= map[string]string{"name":"陈海威2"}
	fmt.Println(dict2["name"])
}

//自动识别字典
func dict03()  {
	dict3:=map[string] string{"name":"chenhaiwei"}
	fmt.Println(dict3["name"])
	dict3["age"]="19"
	fmt.Println(dict3["age"])
}

//make()函数定义字典
func dict04()  {
	var dict4 map[string] string
	dict4=make(map[string] string,4)
	dict4["age"]="10"
	fmt.Println(dict4["age"])
}




