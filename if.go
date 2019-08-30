/*
 if条件判断演示
*/

package main
import "fmt"

func main()  {
	m2("女")
}

//只有一步的条件判断
func m1(i int)  {
	if i>10 {
		fmt.Println("i大于10")
	}
}

//只有两步判断
func m2(k string)  {
	if k=="男" {
		fmt.Print("男")
	}else {
		fmt.Println(k)
	}
}

//有两步以上的条件判断
func m3(j int)  {
	if j==1{
		fmt.Println(j)
	}else if j==2 {
		fmt.Println(j)
	}else {
		fmt.Println("j不是1也不是2")
	}
}












