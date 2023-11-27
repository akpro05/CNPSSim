package failurePage

import (
	"log"

	"github.com/astaxie/beego"
)

type FailurePage struct {
	beego.Controller
}

func (c *FailurePage) Get() {

	log.Println("Debug", "FailurePage request : ")
	defer func() {
		c.TplName = "failurepage/failurepage.html"

		return
	}()

	return
}
