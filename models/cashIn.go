package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

type CashIn struct {
	ID          int     `orm:"auto;ok;column(id)"`
	From        string  `valid:"Required"`
	Amount      float64 `valid:"Required"`
	Description string  `valid:"Required"`
	Details     string  `valid:"Required"`
	Date        int64   `valid:"Required"`
}

type CashInOrmer interface {
	Create(date int64, from string, amount float64, desc string, details string) error
	Read() ([]CashIn, error)
	Update(cashIns *CashIn) error
	Delete(CashInID int) error
}

type BeegoCashInOrmer struct {
	ormer orm.Ormer
}

func NewCashInOrmer(ormer orm.Ormer) CashInOrmer {
	return &BeegoCashInOrmer{ormer: ormer}
}

func (o *BeegoCashInOrmer) Create(date int64, from string, amount float64, desc string, details string) error {
	createCashIns := CashIn{
		Date:        date,
		From:        from,
		Amount:      amount,
		Description: desc,
		Details:     details,
	}

	err := o.ormer.Read(&createCashIns)
	if err != nil {
		return fmt.Errorf("Cant Read At Create, cause: ", err)
	}
	if err != orm.ErrNoRows {
		_, err := o.ormer.Update(&createCashIns)
		if err != nil {
			return fmt.Errorf("Error Updating")
		}
	}
	if err == orm.ErrNoRows {
		_, err := o.ormer.Insert(&createCashIns)
		if err != nil {
			return fmt.Errorf("Error Inserting")
		}
	}
	return nil
}

func (o *BeegoCashInOrmer) Read() ([]CashIn, error) {
	var cashIns []CashIn
	err := o.ormer.Read(&cashIns, "")
	if err == orm.ErrNoRows {
		fmt.Println("Cause of Error: %v", err)
		return nil, nil
	} else if err != nil {
		return nil, fmt.Errorf("Cant read from DB")
	}

	return cashIns, nil
}

func (o *BeegoCashInOrmer) Update(cashIns *CashIn) error {
	_, err := o.ormer.Update(cashIns)
	if err != nil {
		return fmt.Errorf("Cause of Error at Update Cash In: ", err)
	}

	return nil
}

func (o *BeegoCashInOrmer) Delete(cashInID int) error {
	deleteCashIn := CashIn{
		ID: cashInID,
	}

	_, err := o.ormer.Delete(&deleteCashIn)
	if err != nil {
		return fmt.Errorf("Cause of Error at Delete Cash In: ", err)
	}

	return nil
}
