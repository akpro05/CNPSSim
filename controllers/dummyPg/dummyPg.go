package dummyPg

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"

	//	"errors"
	"fmt"
	"io"
	"log"

	"github.com/astaxie/beego"
)

type DummyPg struct {
	beego.Controller
}

func (c *DummyPg) Get() {

	log.Println(beego.AppConfig.String("loglevel"), "Info", "Dummy called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	log.Println("Debug", "Raw Request : ", string(c.GetString("get_pg_ip")))
	defer func() {
		c.TplName = "dummyPg/dummyPg.html"

		return
	}()

	c.Data["indp_declaration_type"] = beego.AppConfig.String("INDP_DECLARATION_TYPE")

	return
}
func (c *DummyPg) Post() {
	//var err error
	log.Println(beego.AppConfig.String("loglevel"), "Info", "Dummy called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	log.Println("Debug", "Raw Request : ", string(c.GetString("get_pg_ip")))
	defer func() {
		c.TplName = "dummyPg/dummyPg.html"
		txn112 := c.Input().Get("cnps_transaction_id")
		log.Println(beego.AppConfig.String("loglevel"), "Debug", "cnps_transaction_id : ", txn112)

		// amount := c.Input().Get("amount")
		// userid := c.Input().Get("user_id")
		// password := beego.AppConfig.String("PGPASS")
		// skey := beego.AppConfig.String("SKEY")

		// cnpssign, _ := checksign(txn, amount, skey, userid, password)
		// if err != nil {
		// 	err = errors.New("Unable to get cnps sign number")
		// 	return
		// }
		// log.Println(beego.AppConfig.String("loglevel"), "Debug", "cnpssign : ", cnpssign)

		return
	}()

	c.Data["indp_declaration_type"] = beego.AppConfig.String("INDP_DECLARATION_TYPE")

	return
}

func checksign(txn, amount, skey, userid, password string) (sign string, err error) {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "txn : ", txn)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "amount : ", amount)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "skey : ", skey)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "userid : ", userid)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "password : ", password)

	input := txn + amount + skey + userid + password
	md5 := md5.New()
	sha_256 := sha256.New()
	sha_512 := sha512.New()
	io.WriteString(md5, input)
	sha_256.Write([]byte(input))
	sha_512.Write([]byte(input))
	sha_512_256 := sha512.Sum512_256([]byte(input))
	hmac512 := hmac.New(sha512.New, []byte(skey))
	hmac512.Write([]byte(input))

	//4db45e622c0ae3157bdcb53e436c96c5
	fmt.Printf("md5:\t\t%x\n", md5.Sum(nil))

	//eb7a03c377c28da97ae97884582e6bd07fa44724af99798b42593355e39f82cb
	fmt.Printf("sha256:\t\t%x\n", sha_256.Sum(nil))

	//5cdaf0d2f162f55ccc04a8639ee490c94f2faeab3ba57d3c50d41930a67b5fa6915a73d6c78048729772390136efed25b11858e7fc0eed1aa7a464163bd44b1c
	fmt.Printf("sha512:\t\t%x\n", sha_512.Sum(nil))

	//34c614af69a2550a4d39138c3756e2cc50b4e5495af3657e5b726c2ac12d5e60
	fmt.Printf("sha512_256:\t%x\n", sha_512_256)

	//GBZ7aqtVzXGdRfdXLHkb0ySp/f+vV9Zo099N+aSv+tTagUWuHrPeECDfUyd5WCoHBe7xkw2EdpyLWx+Ge4JQKg==

	fmt.Printf("hmac512:\t%s\n", base64.StdEncoding.EncodeToString(hmac512.Sum(nil)))
	sign = base64.StdEncoding.EncodeToString(hmac512.Sum(nil))
	return
}
