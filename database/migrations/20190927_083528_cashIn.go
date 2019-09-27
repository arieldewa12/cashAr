package main

import (
	"github.com/astaxie/beego/migration"
)

// DO NOT MODIFY
type CashIn_20190927_083528 struct {
	migration.Migration
}

// DO NOT MODIFY
func init() {
	m := &CashIn_20190927_083528{}
	m.Created = "20190927_083528"

	migration.Register("CashIn_20190927_083528", m)
}

// Run the migrations
func (m *CashIn_20190927_083528) Up() {
	m.SQL(`CREATE TABLE cash_in (
		id int,
		fromTo varchar(225),
		amount double,
		description varchar(225),
		detail varchar(225),
		date varchar(225),
		PRIMARY KEY (id)
		);`)
}

// Reverse the migrations
func (m *CashIn_20190927_083528) Down() {
	// use m.SQL("DROP TABLE ...") to reverse schema update
	m.SQL(`DROP TABLE cash_in;`)
}
