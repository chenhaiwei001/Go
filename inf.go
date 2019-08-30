package main
import "fmt"

//手机开启和关闭
type Phone int
func (p Phone) start() {
	fmt.Println("手机启动中.....")
}
func (p Phone) stop() {
	fmt.Println("手机关闭中.....")
}


//门开启和关闭
type door int
func (d door) start(){
	fmt.Println("门开启中....")
}
func (d door) stop(){
	fmt.Println("门关闭中....")
}

//灯的开启和关闭
type lamp int
func (l lamp) start()  {
	fmt.Println("倒计时5S")
	fmt.Println("电灯开启中......")
}
func (l lamp) stop()  {
	fmt.Println("电灯关闭中......")
}

//窗户的开启和关闭
type window int
func (w window) start()  {
	fmt.Println("窗户开启中......")
}
func (w window) stop()  {
	fmt.Println("窗户关闭中......")
}

//开启和关闭接口
type StartOff interface {
	start()
	stop()
}

func main()  {
   l:=new(window)
   l.start()
   l.stop()
}

















