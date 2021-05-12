package account

import (
	"cilicili-go/conf"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var conn *gorm.DB

func ConnVideoDatabase() {
	var err error
	conn, err = gorm.Open(conf.Config.DB.Type, conf.Config.DB.Url)
	if err != nil {
		panic(err)
	}
	//设置全局表名禁用复数
	conn.SingularTable(true)
}
