/*
   Gorm框架使用示例
*/
package main

import (
	"fmt"
	"net/http"
)

func user(w http.ResponseWriter,request  *http.Request)  {
	query := request.URL.Query()
	id:=query.Get("")
	fmt.Print(id)
}