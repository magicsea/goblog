package main

import (
	"github.com/astaxie/beego"
	_ "github.com/magicsea/goblog/routers"
)

func main() {
	beego.Run()
}
