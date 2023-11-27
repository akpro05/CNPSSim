package statusdummyPg

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"math/rand"
	"time"

	//	"errors"
	"fmt"
	"io"
	"log"

	"github.com/astaxie/beego"
)

type StatusDummyPg struct {
	beego.Controller
}

type Request struct {
	Version                    string `json:"version"`
	Language                   string `json:"language"`
	User_id                    string `json:"user_id"`
	Cnps_transaction_id        string `json:"cnps_transaction_id"`
	Cnps_entity_id             string `json:"cnps_entity_id"`
	Amount                     string `json:"amount"`
	Transaction_end_date       string `json:"transaction_end_date"`
	PaymentMode                string `json:"paymentMode"`
	PaymentMethod              string `json:"paymentMethod"`
	Sign                       string `json:"sign"`
	RequestId                  string `json:"request_id"`
	SuperesbProducerAccessCode string `json:"superesb_producer_access_code"`
}

func (c *StatusDummyPg) Get() {

	log.Println(beego.AppConfig.String("loglevel"), "Info", "Dummy called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	log.Println("Debug", "Raw Request : ", string(c.GetString("get_pg_ip")))

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a 13-digit random number
	randomNumber := generateRandomNumber(13)

	SIMU_1_ACCESS_CODE := beego.AppConfig.String("SIMU_1_ACCESS_CODE")
	c.Data["SIMU_1_ACCESS_CODE"] = SIMU_1_ACCESS_CODE

	c.Data["request_id"] = randomNumber

	defer func() {
		c.TplName = "statusdummyPg/statusdummyPg.html"

		return
	}()

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

func (c *StatusDummyPg) Post() {
	var sreq Request

	log.Println(beego.AppConfig.String("loglevel"), "Info", "status submit called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	defer func() {

		c.TplName = "statusdummyPg/statusdummyPg.html"
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
	PgstatusURL := beego.AppConfig.String("PGSTATUSURL")

	if declaration_type == beego.AppConfig.String("INDP_DECLARATION_TYPE") {
		skey = beego.AppConfig.String("INDPGSKEY")
		password = beego.AppConfig.String("INDPGPASS")
		PgstatusURL = beego.AppConfig.String("INDPGSTATUSURL")
	}

	sign, err := checksign(string(c.GetString("cnps_transaction_id")), string(c.GetString("amount")), string(c.GetString("cnps_entity_id")), skey, string(c.GetString("user_id")), password)

	sreq.Version = string(c.GetString("version"))
	sreq.Language = string(c.GetString("language"))
	sreq.User_id = string(c.GetString("user_id"))
	sreq.Cnps_transaction_id = string(c.GetString("cnps_transaction_id"))
	sreq.Cnps_entity_id = string(c.GetString("cnps_entity_id"))
	sreq.Amount = string(c.GetString("amount"))
	sreq.Transaction_end_date = string(c.GetString("transaction_end_date"))
	sreq.PaymentMode = string(c.GetString("paymentMode"))
	sreq.PaymentMethod = string(c.GetString("paymentMethod"))
	sreq.Sign = sign
	sreq.SuperesbProducerAccessCode = string(c.GetString("superesb_producer_access_code"))
	sreq.RequestId = string(c.GetString("request_id"))

	jsonString, err := json.Marshal(sreq)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", err)
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "jsondata : ", string(jsonString))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", PgstatusURL, bytes.NewBuffer(jsonString)) // URL-encoded payload

	//r, _ := http.NewRequest("POST", "http://167.71.226.158:6002/status/request", bytes.NewBuffer(jsonString)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")

	//fmt.Println(r)

	resp, err := client.Do(r)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Error", err)
		return
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	log.Println("Debug", "Response Body : ", string(body))

	c.Data["response"] = string(body)

	// Generate a 13-digit random number
	randomNumber := generateRandomNumber(13)

	SIMU_1_ACCESS_CODE := beego.AppConfig.String("SIMU_1_ACCESS_CODE")
	c.Data["SIMU_1_ACCESS_CODE"] = SIMU_1_ACCESS_CODE

	c.Data["request_id"] = randomNumber

	return
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
