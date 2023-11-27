package cancelPg

import (
	"log"

	"github.com/astaxie/beego"
)

type CancelPg struct {
	beego.Controller
}

func (c *CancelPg) Get() {

	log.Println("Debug", "---------------This is Cancel Request Received---------------")
	log.Println("Debug", " ")
	log.Println("Debug", "Code : ", c.GetString("Code"))
	log.Println("Debug", "Message : ", c.GetString("Message"))
	log.Println("Debug", "cnps_transaction_id : ", c.GetString("cnps_transaction_id"))
	log.Println("Debug", " ")
	log.Println("Debug", "---------------This is Cancel Request Finished---------------")

	c.Ctx.Output.Body([]byte("This is cancel page \n\n" + "code :" + c.GetString("Code") + "\nMessage :" + c.GetString("Message") + "\ncnps_transaction_id :" + c.GetString("cnps_transaction_id")))

	return
}
