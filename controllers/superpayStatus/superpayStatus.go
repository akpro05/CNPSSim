package superpayStatus

import (
	"bytes"

	"encoding/json"
	"io/ioutil"
	"net/http"

	"math/rand"
	"time"

	//	"errors"
	"fmt"

	"log"

	"github.com/astaxie/beego"
)

type SuperpayStatus struct {
	beego.Controller
}

type Request struct {
	AccessCode                 string `json:"access_code"`
	SuperesbProducerAccessCode string `json:"superesb_producer_access_code"`
	Channel                    string `json:"channel"`
	TxnNumber                  string `json:"txn_number"`
	Mobile                     string `json:"mobile"`
	RequestID                  string `json:"request_id"`
	Language                   string `json:"language"`
	UserType                   string `json:"user_type"`
}

func (c *SuperpayStatus) Get() {

	log.Println(beego.AppConfig.String("loglevel"), "Info", "SuperpayStatus called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	log.Println("Debug", "Raw Request : ", string(c.GetString("get_pg_ip")))

	// Seed the random number generator with the current time
	rand.Seed(time.Now().UnixNano())

	// Generate a 13-digit random number
	randomNumber := generateRandomNumber(13)

	c.Data["request_id"] = randomNumber

	SIMU_1_ACCESS_CODE := beego.AppConfig.String("SIMU_1_ACCESS_CODE")
	c.Data["SIMU_1_ACCESS_CODE"] = SIMU_1_ACCESS_CODE

	defer func() {
		c.TplName = "superpayStatus/superpayStatus.html"

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

func (c *SuperpayStatus) Post() {
	var sreq Request

	log.Println(beego.AppConfig.String("loglevel"), "Info", "SuperpayStatus called")

	log.Println("Debug", "Raw Request : ", string(c.Ctx.Input.RequestBody))

	defer func() {
		c.TplName = "superpayStatus/superpayStatus.html"
		return
	}()

	SPAY_GET_STATUS_URL := beego.AppConfig.String("SPAY_GET_STATUS_URL")

	sreq.AccessCode = string(c.GetString("access_code"))
	sreq.SuperesbProducerAccessCode = string(c.GetString("superesb_producer_access_code"))
	sreq.Channel = string(c.GetString("channel"))
	sreq.TxnNumber = string(c.GetString("txn_number"))
	sreq.Mobile = string(c.GetString("mobile"))
	sreq.RequestID = string(c.GetString("request_id"))
	sreq.Language = string(c.GetString("language"))
	sreq.UserType = string(c.GetString("user_type"))

	jsonString, err := json.Marshal(sreq)
	if err != nil {
		log.Println(beego.AppConfig.String("loglevel"), "Debug", err)
		return
	}

	log.Println(beego.AppConfig.String("loglevel"), "Debug", "jsondata : ", string(jsonString))

	client := &http.Client{}
	r, _ := http.NewRequest("POST", SPAY_GET_STATUS_URL, bytes.NewBuffer(jsonString)) // URL-encoded payload

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
	randomNumber := generateRandomNumber(13)

	c.Data["request_id"] = randomNumber

	SIMU_1_ACCESS_CODE := beego.AppConfig.String("SIMU_1_ACCESS_CODE")
	c.Data["SIMU_1_ACCESS_CODE"] = SIMU_1_ACCESS_CODE

	return
}
