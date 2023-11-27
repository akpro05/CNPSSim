package returnPg

import (
	"log"

	"github.com/astaxie/beego"
)

type ReturnPg struct {
	beego.Controller
}

func (c *ReturnPg) Get() {

	log.Println("Debug", "---------------This is Return Request Received---------------")
	log.Println("Debug", " ")
	log.Println("Debug", "cnps_transaction_id : ", c.GetString("cnps_transaction_id"))
	log.Println("Debug", "pgs_transaction_id : ", c.GetString("pgs_transaction_id"))
	log.Println("Debug", "amount : ", c.GetString("amount"))
	log.Println("Debug", "code : ", c.GetString("code"))
	log.Println("Debug", "status : ", c.GetString("status"))
	log.Println("Debug", "pgs_transaction_date : ", c.GetString("pgs_transaction_date"))

	log.Println("Debug", " ")
	log.Println("Debug", "---------------This is Return Request Finished---------------")

	c.Ctx.Output.Body([]byte("This is Return page \n\n" + "code :" + c.GetString("code") +
		"\npgs_transaction_id :" + c.GetString("pgs_transaction_id") +
		"\namount :" + c.GetString("amount") +
		"\nstatus :" + c.GetString("status") +
		"\npgs_transaction_date :" + c.GetString("pgs_transaction_date") +
		"\ncnps_transaction_id :" + c.GetString("cnps_transaction_id")))

	return
}
