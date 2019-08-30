package main

import "fmt"


func mux2(c *int) {
	 *c=19
}

func main()  {
	var a int=10
	var b *int=&a
	fmt.Println("更改前：",a)
	mux2(b)
	fmt.Println("更改后：",a)
}