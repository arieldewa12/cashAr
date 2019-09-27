package registration

import (
	"fmt"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"cashAr/models"

	"github.com/astaxie/beego/orm"
)

type databaseConfig struct {
	Username string
	Password string
	Host     string
	Name     string
}

func init() {
	max_conn := 30
	max_idle_conn := 30
	max_timeout := 3
	databaseConfig := databaseConfig{
		Username: "root",
		Password: "root",
		Host:     "db",
		Name:     "finance",
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)

	connString := fmt.Sprintf("mysql://%s:%s@%s/%s?connect_timeout=%s", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Name, strconv.Itoa(max_timeout))
	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDataBase("default", "mysql", connString, max_idle_conn, max_conn)

	//Register all models here
	orm.RegisterModel(new(models.CashIn))
	orm.RegisterModel(new(models.CashOut))
}
