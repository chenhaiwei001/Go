/*
   switch两种条件判断演示
*/
package main
import "fmt"
func main()  {
	s2(5)
}

//switch条件判断,针对变量的不同值作判断
func s1(i int)  {
	switch i{
	case 0:
		fmt.Println("零")
	case 1:
		fmt.Println("一")
	case 2:
		fmt.Println("二")
	default:
		fmt.Println("不知道什么值")
	}
}


//每个判断单独一个表达式
func s2(i int)  {
	switch{
	    case i==1:
	    	fmt.Println("一年级")
	    case i==2:
	     	fmt.Println("二年级")
	case i==3:
		    fmt.Println("三年级")
	default:
		    fmt.Println("大于三年级")
	}
}














