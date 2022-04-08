package main

//用于练习golang基本语法和简单的功能实现.
/**
* @Author: 周佩瑾
* @Description: 演示接口
* @File: main.go
* @Date: 2022/3/21 17:06
 */
import (
	"demo/routerfunc"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	// var a int = 3
	// var param interface{}
	// param = a
	// fmt.Println("param:",param.(int))
	mus := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mus.Handle("/", fs)

	router := mux.NewRouter()
	router.HandleFunc("/query/{id}", routerfunc.SearchPeopleWithID).Methods("GET")    //查找
	router.HandleFunc("/insert", routerfunc.InsertPeople).Methods("POST")             //插入
	router.HandleFunc("/updata/{id}", routerfunc.UpdatePeople).Methods("POST")        //更新
	router.HandleFunc("/delete/{id}", routerfunc.DeletePeopleUseID).Methods("DELETE") //删除

	fmt.Println("Server at localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", router)) //建立服务器
}
