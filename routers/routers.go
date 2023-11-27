package routers

import (
	"CNPSSim/controllers/cancelPg"
	"CNPSSim/controllers/dummyPg"
	"CNPSSim/controllers/failurepage"
	"CNPSSim/controllers/home"
	"CNPSSim/controllers/notificationPg"
	"CNPSSim/controllers/returnPg"
	"CNPSSim/controllers/statusdummyPg"
	"CNPSSim/controllers/submitPg"
	"CNPSSim/controllers/successpage"
	"CNPSSim/controllers/superpayStatus"

	"github.com/astaxie/beego"
)

func init() {

	beego.Router(beego.AppConfig.String("HOME_PATH"), &home.Home{})
	beego.Router(beego.AppConfig.String("RETURN_PATH"), &returnPg.ReturnPg{})
	beego.Router(beego.AppConfig.String("CANCEL_PATH"), &cancelPg.CancelPg{})
	beego.Router(beego.AppConfig.String("NOTIFICATION_PATH"), &notificationPg.NotifyPg{})
	beego.Router(beego.AppConfig.String("DUMMY_PATH"), &dummyPg.DummyPg{})
	beego.Router(beego.AppConfig.String("SUBMIT_PATH"), &submitPg.SubmitPg{})

	beego.Router(beego.AppConfig.String("SUCCESSPAGE"), &successPage.SuccessPage{})
	beego.Router(beego.AppConfig.String("FAILUREPAGE"), &failurePage.FailurePage{})

	beego.Router(beego.AppConfig.String("STATUS_DUMMY_PATH"), &statusdummyPg.StatusDummyPg{})

	beego.Router(beego.AppConfig.String("SPAY_TXN_STATUS_PATH"), &superpayStatus.SuperpayStatus{})

}
