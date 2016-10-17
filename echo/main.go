package main

import (
	"com.izuanqian/echo/controller"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func main() {
	m := martini.Classic()
	m.Use(render.Renderer())
	controller.Api(m)
	m.RunOnAddr(":80")

}
