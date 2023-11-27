package successPage

import (
	"log"

	"github.com/astaxie/beego"
)

type SuccessPage struct {
	beego.Controller
}

func (c *SuccessPage) Get() {

	log.Println("Debug", "SuccessPage request : ")
	defer func() {
		c.TplName = "successpage/successpage.html"

		return
	}()

	return
}
