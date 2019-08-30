/*
 指针的使用说明：
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
