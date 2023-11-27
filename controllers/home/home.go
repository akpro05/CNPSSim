package home

import (
	"log"

	"github.com/astaxie/beego"
)

type Home struct {
	beego.Controller
}

func (c *Home) Get() {

	log.Println("Debug", "Home request : ")
	defer func() {
		c.TplName = "home/home.html"

		return
	}()

	return
}
