/*
   常用的几种for循环演示
*/

package main
import "fmt"

func main()  {
	k1()
}

//从0到100循环输出（标准格式）
func k1()  {
	for i:=0;i<=100;i++{
		println(i)
	}
}


//只有一个条件判断的循环
func k2()  {
	var i int=10
	var j int=1
	for i>=j  {
		fmt.Println(j)
		j++
	}
}


//死循环（不建议使用）
func k3()  {
	var i int
	for{
		 i++
		fmt.Println("死循环+",i)
	}
}




