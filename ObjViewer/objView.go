package main

import (
	"com.izuanqian/ObjViewer/handler"
)

type eventHandler interface {
	Do()
}

func main() {

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
