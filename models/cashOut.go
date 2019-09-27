package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

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
	Create(client string, invoice float64, pv string, accNum string, to string, bank string, accNumTo string, amount float64, bnkCharge float64, journal string, doc string, dpp float64, ppn float64, taxInvNum int64, taxInvDate string, company string) error
	Read() ([]CashOut, error)
	Update(cashOut *CashOut) error
	Delete(cashOutID int) error
}

type BeegoCashOutOrmer struct {
	ormer orm.Ormer
}

func NewCashOutOrmer(ormer orm.Ormer) CashOutOrmer {
	return &BeegoCashOutOrmer{ormer: ormer}
}

func (o *BeegoCashOutOrmer) Create(client string, invoice float64, pv string, accNum string, to string, bank string, accNumTo string, amount float64, bnkCharge float64, journal string, doc string, dpp float64, ppn float64, taxInvNum int64, taxInvDate string, company string) error {
	createCashOuts := CashOut{
		Client:           client,
		Invoice:          invoice,
		Pv:               pv,
		AccountNumber:    accNum,
		To:               to,
		Bank:             bank,
		AccountNumberTo:  accNumTo,
		Amount:           amount,
		BankCharge:       bnkCharge,
		Journal:          journal,
		Document:         doc,
		Dpp:              dpp,
		Ppn:              ppn,
		TaxInvoiceNumber: taxInvNum,
		TaxInvoiceDate:   taxInvDate,
		CompanyName:      company,
	}

	err := o.ormer.Read(&createCashOuts)
	if err != nil {
		return fmt.Errorf("Cant Read At Create, cause: ", err)
	}
	if err != orm.ErrNoRows {
		_, err := o.ormer.Update(&createCashOuts)
		if err != nil {
			return fmt.Errorf("Error Updating")
		}
	}
	if err == orm.ErrNoRows {
		_, err := o.ormer.Insert(&createCashOuts)
		if err != nil {
			return fmt.Errorf("Error Inserting")
		}
	}
	return nil
}

func (o *BeegoCashOutOrmer) Read() ([]CashOut, error) {
	var cashOuts []CashOut
	err := o.ormer.Read(&cashOuts, "")
	if err == orm.ErrNoRows {
		fmt.Println("Cause of Error: %v", err)
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("Cant read from DB")
	}

	return cashOuts, nil
}

func (o *BeegoCashOutOrmer) Update(cashOut *CashOut) error {
	_, err := o.ormer.Update(cashOut)
	if err != nil {
		return fmt.Errorf("Cause of Error at Update Cash Out: ", err)
	}
	return nil
}

func (o *BeegoCashOutOrmer) Delete(cashOutID int) error {
	deleteCashOut := CashOut{
		ID: cashOutID,
	}

	_, err := o.ormer.Delete(&deleteCashOut)
	if err != nil {
		return fmt.Errorf("Cause of Error at Delete Cash In: ", err)
	}

	return nil
}
