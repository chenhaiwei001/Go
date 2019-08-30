/*
  方法的定义及使用演示
*/
package main

import "fmt"

//学生的跑和跳
type student string
//跑的方法
func (s student) run()  {
	fmt.Println("跑")
}
//跳的方法
func (s student) dump()  {
	fmt.Println("跳")
}

//教师的跑和跳
type teacher string

//跑的方法
func (t teacher) run() {
	fmt.Println("跑")
}
//跳
func (t teacher) dump() {
	fmt.Println("跳")
}


func main()  {
	var s student  //实例化一个方法
	s.dump()     //调用dump方法
	s.run()    //调用run方法

	var t teacher
	t.run()
	t.dump()

}