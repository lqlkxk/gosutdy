package main

import (
	"html/template"
	"log"
	"net/http"
)


func main(){
	listen()
}

func sayHello(response http.ResponseWriter ,request *http.Request){

	temp,err := template.ParseFiles("./template/hello.tmpl")
	if err !=nil{
		log.Panicf("模板渲染失败:%s/n",err)
		panic(err)
	}
	temp.Execute(response,1)
}

func listen(){
	http.HandleFunc("/",sayHello)
	err := http.ListenAndServe(":8080",nil)
	if err !=nil{
		log.Panicf("模板渲染失败:%s/n",err)
		panic(err)
	}

}