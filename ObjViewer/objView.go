package main

import (
	"com.izuanqian/ObjViewer/db"
	"com.izuanqian/ObjViewer/handler"
	"encoding/json"
	"fmt"
)

type eventHandler interface {
	Do()
}

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

func main() {
	var user User
	WrapViewObj("USER", "0001", &user)
	fmt.Println(user)
	doHandler()
}

func doHandler() {
	c := make(chan int, 10)
	{
		go handler.BLPOPA()
		go handler.BLPOPB()
	}
	<-c
}

// 从redis加载object
// 将对象赋值到viewObj
func WrapViewObj(key string, objSymbol string, viewObj interface{}) {

	reply := db.GetOptConnection().Cmd("GET", key+":"+objSymbol)
	objJson, _ := reply.Str()
	if len(objJson) != 0 {
		err := json.Unmarshal([]byte(objJson), viewObj)
		if err != nil {
			fmt.Println(err)
		}

	}
}
