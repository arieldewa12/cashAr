package models

type CashIn struct {
	ID          int     `orm:"auto;ok;column(id)"`
	From        string  `valid:"Required"`
	Amount      float64 `valid:"Required"`
	Description string  `valid:"Required"`
	Details     string  `valid:"Required"`
	Date        int64   `valid:"Required"`
}
