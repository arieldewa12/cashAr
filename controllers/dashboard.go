package controllers

import (
	"github.com/astaxie/beego"
	// "github.com/astaxie/beego/orm"
)

type DashboardController struct {
	beego.Controller
	// cashInOrmer  models.CashInOrmer
	// cashOutOrmer models.CashOutOrmer
}

type DashboardContent struct {
	CashIn Currency
	CashOut Currency
	Ar Currency 
	Ap Currency
}

type Currency struct {
	Value int 
	Currency string
}

func (ths DashboardController) Prepare() {
	
}

func (ths *DashboardController) Get() {

	ths.Data["content"] = DashboardContent{
		CashIn : Currency{
			Value : 5000,
			Currency : "IDR"},
		CashOut : Currency{
			Value : 5000,
			Currency : "IDR"},
		Ar : Currency{
			Value : 5000,
			Currency : "IDR"},
		Ap : Currency{
			Value : 5000,
			Currency : "IDR"},
	}
	ths.TplName = "dashboard.html"
}

// func (ths *DashboardController) Post() {

// }
