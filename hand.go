/*
    指针的使用说明：
    这是没有使用指针的时候，调用同一变量的情况
*/
package main

import "fmt"

func mux(j int) {
     j=11
}


func main()  {
	var i int =10
	fmt.Println("更改前：",i)
	mux(i)
	fmt.Println("更改后：",i)
}
