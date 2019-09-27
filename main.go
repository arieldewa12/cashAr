package main

import (
	_ "cashAr/routers"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
