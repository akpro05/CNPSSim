package submitPg

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"io"

	"math/rand"
	"time"

	//	"errors"
	"log"

	"github.com/astaxie/beego"
)

type SubmitPg struct {
	beego.Controller
}

func (c *SubmitPg) Post() {

	log.Println(beego.AppConfig.String("loglevel"), "Info", "submit called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	defer func() {
		c.TplName = "submitPg/submitPg.html"

		return
	}()

	declaration_type := c.GetString("declaration_type")

	if declaration_type == "" {
		log.Panic("Please select declaration type")
		return
	}
	fmt.Println(beego.AppConfig.String("loglevel"), "Debug", "DeclarationType : ", declaration_type)

	skey := beego.AppConfig.String("SKEY")
	password := beego.AppConfig.String("PGPASS")
	pgs_url := beego.AppConfig.String("EMPPGSURL")

	if declaration_type == beego.AppConfig.String("INDP_DECLARATION_TYPE") {
		skey = beego.AppConfig.String("INDPGSKEY")
		password = beego.AppConfig.String("INDPGPASS")
		pgs_url = beego.AppConfig.String("INDPGSURL")
	}

	sign, err := checksign(string(c.GetString("cnps_transaction_id")), string(c.GetString("amount")), string(c.GetString("cnps_entity_id")), skey, string(c.GetString("user_id")), password)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "error : ", err)

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "sdasdd : ", string(c.GetString("cnps_entity_id")))

	c.Data["cnps_entity_id"] = string(c.GetString("cnps_entity_id"))
	c.Data["entity_name"] = string(c.GetString("entity_name"))
	c.Data["language"] = string(c.GetString("language"))
	c.Data["entity_phone"] = string(c.GetString("entity_phone"))
	c.Data["entity_email"] = string(c.GetString("entity_email"))
	c.Data["amount"] = string(c.GetString("amount"))
	c.Data["currency"] = string(c.GetString("currency"))
	c.Data["cnps_transaction_id"] = string(c.GetString("cnps_transaction_id"))
	c.Data["partial_payment"] = string(c.GetString("partial_payment"))
	c.Data["transaction_start_date"] = string(c.GetString("transaction_start_date"))
	c.Data["notification_url"] = string(c.GetString("notification_url"))
	c.Data["return_url"] = string(c.GetString("return_url"))
	c.Data["cancel_url"] = string(c.GetString("cancel_url"))
	c.Data["description"] = string(c.GetString("description"))
	c.Data["user_id"] = string(c.GetString("user_id"))
	c.Data["declaration_type"] = string(c.GetString("declaration_type"))
	c.Data["nature_code"] = string(c.GetString("nature_code"))
	c.Data["nature_name"] = string(c.GetString("nature_name"))
	c.Data["declaration_period"] = string(c.GetString("declaration_period"))
	c.Data["cnps_declaration_number"] = string(c.GetString("cnps_declaration_number"))
	c.Data["version"] = string(c.GetString("version"))
	c.Data["pgs_url"] = pgs_url
	c.Data["indp_declaration_type"] = beego.AppConfig.String("INDP_DECLARATION_TYPE")
	c.Data["payment_mode"] = string(c.GetString("payment_mode"))
	c.Data["payment_method"] = string(c.GetString("payment_method"))
	c.Data["customer_mobile_number"] = string(c.GetString("customer_mobile_number"))
	c.Data["payment_type"] = string(c.GetString("payment_type"))
	SIMU_1_ACCESS_CODE := beego.AppConfig.String("SIMU_1_ACCESS_CODE")
	c.Data["SIMU_1_ACCESS_CODE"] = SIMU_1_ACCESS_CODE

	c.Data["sign"] = sign

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a 13-digit random number
	randomNumber := generateRandomNumber(13)

	c.Data["request_id"] = randomNumber

	return
}

func generateRandomNumber(digits int) string {
	min := int64(1e12) // 1 followed by 12 zeros
	max := int64(1e13) // 1 followed by 13 zeros

	// Generate a random number within the specified range
	randomValue := rand.Int63n(max-min) + min

	// Convert the random number to a string with leading zeros if necessary
	randomNumber := fmt.Sprintf("%0*d", digits, randomValue)

	return randomNumber
}

func checksign(txn, amount, cnps_entity_id, skey, userid, password string) (sign string, err error) {
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "txn : ", txn)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "amount : ", amount)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "cnps_entity_id : ", cnps_entity_id)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "skey : ", skey)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "userid : ", userid)
	log.Println(beego.AppConfig.String("loglevel"), "Debug", "password : ", password)

	input := txn + amount + cnps_entity_id + skey + userid + password
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
