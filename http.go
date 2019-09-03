package main

import (
	"fmt"
	"net/http"
	"time"
)

//获取get参数
func hello(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query() //实例化对象
	name := query.Get("name")    //获取get参数
	fmt.Println(time.Now(),"成功获取GET参数:",name)
	fmt.Fprintf(w,name)         //输出
}

//获取POST参数
func word(w http.ResponseWriter, request *http.Request) {
	name := request.PostFormValue("name")//获取post参数
	fmt.Fprintf(w, name)
	fmt.Println(time.Now(),"成功获取POST参数:", name)
}

//主函数
func main() {
	fmt.Println("后台等待中...")
	//设置端口和ip地址
	server := &http.Server{
		Addr: ":80",
		Handler:nil,
	}
	http.HandleFunc("/hello", hello) //路由转发到hello函数
	http.HandleFunc("/word", word)   //路由转发至word函数
	server.ListenAndServe()          //监听端口
}
