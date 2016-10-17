package main

import (
	"com.izuanqian/ObjViewer/domain"
	"com.izuanqian/ObjViewer/handler"
	"encoding/json"
	"fmt"
	"log"
)

type eventHandler interface {
	Do()
}

func main() {

	v, _ := json.Marshal(domain.OrderCreate{Id: "00010", Price: 100})
	fmt.Println(string(v))
	log.Println("handle running...")
	doHandler()

}

func doHandler() {
	//go func() {
	for {
		//handler.OrderCreateHandler{}.Do()
		handler.BLPOP()
	}
	//}()
}
