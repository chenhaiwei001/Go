/*
   指针的使用示例
*/

package main
import "fmt"

//定义一个指针函数
func mux2(c *int) {
	 *c=19
}

//主函数
func main()  {
	var a int=10   //声明一个变量
	var b *int=&a  //声明一个指针
	fmt.Println("更改前：",a)
	mux2(b)   //对变量的值更改后
	fmt.Println("更改后：",a)
}
/*
   在函数内更改变量后，原始变量的值也更改了，说明使用指针可以更改原始变量的值
   应用场景：比如需要与另一个函数进行数据交互的情况
*/