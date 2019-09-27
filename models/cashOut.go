package models

import "github.com/astaxie/beego/orm"

type CashOut struct {
	ID               int     `orm:"auto;ok;column(id)"`
	Client           string  `valid:"Required"`
	Invoice          float64 `valid:"Required"`
	Pv               string  `valid:"Required"`
	AccountNumber    string  `valid:"Required"`
	To               string  `valid:"Required"`
	Bank             string  `valid:"Required"`
	AccountNumberTo  string  `valid:"Required"`
	Amount           float64 `valid:"Required"`
	BankCharge       float64 `valid:"Required"`
	Journal          string  `valid:"Required"`
	Document         string  `valid:"Required"`
	Dpp              float64 `valid:"Required"`
	Ppn              float64 `valid:"Required"`
	TaxInvoiceNumber int64   `valid:"Required"`
	TaxInvoiceDate   string  `valid:"Required"`
	CompanyName      string  `valid:"Required"`
}

type CashOutOrmer interface {
	Create() error
	Read() error
	Update() error
	Delete() error
	AddCashOutFromCSV(CashOut []CashOut) error
}

type BeegoCashOutOrmer struct {
	ormer orm.Ormer
}

func NewCashOutOrmer(ormer orm.Ormer) CashOutOrmer {
	return &BeegoCashOutOrmer{ormer: ormer}
}

func (c *BeegoCashOutOrmer) Create() error {
	return nil
}

func (c *BeegoCashOutOrmer) Read() error {
	return nil
}

func (c *BeegoCashOutOrmer) Update() error {
	return nil
}

func (C *BeegoCashOutOrmer) Delete() error {
	return nil
}

func (o *BeegoCashOutOrmer) AddCashOutFromCSV(CashOut []CashOut) error {
	for _, value := range CashOut {
		err := o.Create(value.Date, "", value.Amount, value.Description, "")
		if err != nil {
			return err
		}
	}
	return nil
}
