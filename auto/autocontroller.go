package auto

import (
	"./controller"
	"../service/controller"
	"../service/data"
	"log"
	"net/http"
)

func autoAuthenticate(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("before process") // 処理の前の共通処理
		user := data.GetStringSession(r, "user")
		if user != ""{
			fn(w, r)
		} else {
			controller.GetLoginViewHandler(w, r)
		}

		log.Println("after process") // 処理の後の共通処理
	}
}

func AutoControllerLoad() {
	http.HandleFunc("/autosample", autoAuthenticate(autocontroller.AutoIndexViewHandler))

	http.HandleFunc("/autosample/HMailLog", autoAuthenticate(autocontroller.HMailLogSearchViewHandler))
	http.HandleFunc("/autosample/HMailLog/register", autoAuthenticate(autocontroller.HMailLogRegisterViewHandler))
	http.HandleFunc("/autosample/MBlock", autoAuthenticate(autocontroller.MBlockSearchViewHandler))
	http.HandleFunc("/autosample/MBlock/register", autoAuthenticate(autocontroller.MBlockRegisterViewHandler))
	http.HandleFunc("/autosample/MIshiguro", autoAuthenticate(autocontroller.MIshiguroSearchViewHandler))
	http.HandleFunc("/autosample/MIshiguro/register", autoAuthenticate(autocontroller.MIshiguroRegisterViewHandler))
	http.HandleFunc("/autosample/MSetting", autoAuthenticate(autocontroller.MSettingSearchViewHandler))
	http.HandleFunc("/autosample/MSetting/register", autoAuthenticate(autocontroller.MSettingRegisterViewHandler))
	http.HandleFunc("/autosample/MSettingConnectionMessage", autoAuthenticate(autocontroller.MSettingConnectionMessageSearchViewHandler))
	http.HandleFunc("/autosample/MSettingConnectionMessage/register", autoAuthenticate(autocontroller.MSettingConnectionMessageRegisterViewHandler))
	http.HandleFunc("/autosample/MShop", autoAuthenticate(autocontroller.MShopSearchViewHandler))
	http.HandleFunc("/autosample/MShop/register", autoAuthenticate(autocontroller.MShopRegisterViewHandler))
	http.HandleFunc("/autosample/MUser", autoAuthenticate(autocontroller.MUserSearchViewHandler))
	http.HandleFunc("/autosample/MUser/register", autoAuthenticate(autocontroller.MUserRegisterViewHandler))
	http.HandleFunc("/autosample/RFirstConversion", autoAuthenticate(autocontroller.RFirstConversionSearchViewHandler))
	http.HandleFunc("/autosample/RFirstConversion/register", autoAuthenticate(autocontroller.RFirstConversionRegisterViewHandler))
	http.HandleFunc("/autosample/TBeaconAccess", autoAuthenticate(autocontroller.TBeaconAccessSearchViewHandler))
	http.HandleFunc("/autosample/TBeaconAccess/register", autoAuthenticate(autocontroller.TBeaconAccessRegisterViewHandler))
	http.HandleFunc("/autosample/TCompareAccess", autoAuthenticate(autocontroller.TCompareAccessSearchViewHandler))
	http.HandleFunc("/autosample/TCompareAccess/register", autoAuthenticate(autocontroller.TCompareAccessRegisterViewHandler))
	http.HandleFunc("/autosample/TInquiry", autoAuthenticate(autocontroller.TInquirySearchViewHandler))
	http.HandleFunc("/autosample/TInquiry/register", autoAuthenticate(autocontroller.TInquiryRegisterViewHandler))
	http.HandleFunc("/autosample/TMemo", autoAuthenticate(autocontroller.TMemoSearchViewHandler))
	http.HandleFunc("/autosample/TMemo/register", autoAuthenticate(autocontroller.TMemoRegisterViewHandler))
	http.HandleFunc("/autosample/TSettingConnection", autoAuthenticate(autocontroller.TSettingConnectionSearchViewHandler))
	http.HandleFunc("/autosample/TSettingConnection/register", autoAuthenticate(autocontroller.TSettingConnectionRegisterViewHandler))
	http.HandleFunc("/autosample/TTest", autoAuthenticate(autocontroller.TTestSearchViewHandler))
	http.HandleFunc("/autosample/TTest/register", autoAuthenticate(autocontroller.TTestRegisterViewHandler))
	http.HandleFunc("/autosample/TTest2", autoAuthenticate(autocontroller.TTest2SearchViewHandler))
	http.HandleFunc("/autosample/TTest2/register", autoAuthenticate(autocontroller.TTest2RegisterViewHandler))

}
