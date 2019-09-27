package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CashOut_20190927_101531 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CashOut_20190927_101531{}
	m.Created = "20190927_101531"

	migration.Register("CashOut_20190927_101531", m)
}

// Run the migrations
func (m *CashOut_20190927_101531) Up() {
	// use m.SQL("CREATE TABLE ...") to make schema update
	m.SQL(`CREATE TABLE cash_out (
		id int, 
		client varchar(225),
		invoice varchar(225),
		pv varchar(225),
		account_number varchar(225),
		toFrom varchar(225),
		bank varchar(225),
		account_number_to varchar(225),
		amount double,
		bank_charge double,
		journal varchar(225),
		documents varchar(225),
		dpp_double double,
		ppn_double double,
		tax_invoice_number double,
		tax_invoice_date varchar(225),
		company_name varchar(225),
		PRIMARY KEY (id)
	);`)
}

// Reverse the migrations
func (m *CashOut_20190927_101531) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE cash_out;`)
}
