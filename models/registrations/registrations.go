package registration

import (
	"cashAr/models"
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type databaseConfig struct {
	Username string
	Password string
	Host     string
	Name     string
}

func init() {
	// max_conn := 30
	// max_idle_conn := 30
	// max_timeout := 3
	databaseConfig := databaseConfig{
		Username: "root",
		Password: "",
		Host:     "mysql",
		Name:     "root",
	}

	connString := fmt.Sprintf(`%s:%s@/%s?charset=utf8`, databaseConfig.Username, databaseConfig.Password, databaseConfig.Name)
	//connString := fmt.Sprintf("mysql://%s:%s@%s/%s?connect_timeout=%s&sslmode=disable", databaseConfig.Username, databaseConfig.Password, databaseConfig.Host, databaseConfig.Name, strconv.Itoa(max_timeout))

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.DefaultTimeLoc = time.UTC

	orm.RegisterDataBase("default", "mysql", connString)

	//Register all models here
	orm.RegisterModel(new(models.CashIn))
	orm.RegisterModel(new(models.CashOut))
}
