package notificationPg

import (
	"log"

	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"

	"github.com/astaxie/beego"
)

type NotifyPg struct {
	beego.Controller
}

func (c *NotifyPg) Post() {
	request := c.Ctx.Request
	log.Println("Debug", "Debug", "Notify Get Request : ", request)
	pgs_transaction_id := c.GetString("pgs_transaction_id")
	amount := c.GetString("amount")
	status := c.GetString("status")
	paymentMode := c.GetString("paymentMode")
	cnps_transaction_id := c.GetString("cnps_transaction_id")
	log.Println("Debug", "---------------This is Return Notification Received---------------")
	log.Println("Debug", " ")
	log.Println("Debug", "cnps_transaction_id : ", cnps_transaction_id)
	log.Println("Debug", "partial_payment : ", c.GetString("partial_payment"))
	log.Println("Debug", "transaction_end_date : ", c.GetString("transaction_end_date"))
	log.Println("Debug", "pgs_transaction_id : ", pgs_transaction_id)
	log.Println("Debug", "code : ", c.GetString("code"))
	log.Println("Debug", "status : ", status)
	log.Println("Debug", "message : ", c.GetString("message"))
	log.Println("Debug", "pgs_transaction_date : ", c.GetString("pgs_transaction_date"))
	log.Println("Debug", "amount : ", amount)
	log.Println("Debug", "currecncy : ", c.GetString("currecncy"))
	log.Println("Debug", "paymentMethod : ", c.GetString("paymentMethod"))
	log.Println("Debug", "paymentMode : ", paymentMode)
	log.Println("Debug", "Sign : ", c.GetString("sign"))

	// log.Println("Debug", "code : ", c.GetString("code"))
	// log.Println("Debug", "Message : ", c.GetString("message"))
	// log.Println("Debug", "cnps_transaction_id : ", c.GetString("cnps_transaction_id"))
	log.Println("Debug", " ")
	log.Println("Debug", "---------------This is Return Notification Finished---------------")
	INDPGSKEY := beego.AppConfig.String("INDPGSKEY")
	signParam := pgs_transaction_id + amount + status + cnps_transaction_id + paymentMode + INDPGSKEY + "TEST001"
	log.Println("Debug", "Sign Params  : ", signParam)
	log.Println("Debug", "Generated Sign : ", makesignofParam(signParam, INDPGSKEY))
	c.Ctx.Output.Body([]byte("This is Notification page \n\n" + "code :" + c.GetString("code") + "\nMessage :" + c.GetString("message") + "\ncnps_transaction_id :" + c.GetString("cnps_transaction_id")))

}

func makesignofParam(InputString, skey string) (sign string) {

	input := InputString
	hmac512 := hmac.New(sha512.New, []byte(skey))
	hmac512.Write([]byte(input))

	//4db45e622c0ae3157bdcb53e436c96c5
	//fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))

	//eb7a03c377c28da97ae97884582e6bd07fa44724af99798b42593355e39f82cb
	//fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))

	//5cdaf0d2f162f55ccc04a8639ee490c94f2faeab3ba57d3c50d41930a67b5fa6915a73d6c78048729772390136efed25b11858e7fc0eed1aa7a464163bd44b1c
	//fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))

	//34c614af69a2550a4d39138c3756e2cc50b4e5495af3657e5b726c2ac12d5e60
	//fmt.Printf("sha512_256:\t%x\n", sha_512_256)

	//GBZ7aqtVzXGdRfdXLHkb0ySp/f+vV9Zo099N+aSv+tTagUWuHrPeECDfUyd5WCoHBe7xkw2EdpyLWx+Ge4JQKg==

	sign = base64.StdEncoding.EncodeToString(hmac512.Sum(nil))
	return
}
