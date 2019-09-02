package main

import (
	"fmt"
	"net/http"
)

//GET请求实例
func hello(w http.ResponseWriter, request *http.Request) {
	query := request.URL.Query() //实例化对象
	name := query.Get("name")    //获取get参数
	fmt.Fprintf(w, name)         //输出
}

//post请求实例
func word(w http.ResponseWriter, request *http.Request) {
	name := request.PostFormValue("name")//获取post参数
	fmt.Fprintf(w, name)
}

//主函数
func main() {
	//设置端口和ip地址
	server := &http.Server{
		Addr: "0.0.0.0:80",
	}
	http.HandleFunc("/hello", hello) //路由转发到hello函数
	http.HandleFunc("/word", word)   //路由转发至word函数
	server.ListenAndServe()          //监听端口
}
