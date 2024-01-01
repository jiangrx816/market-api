package initialize

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"wechat/global"
)

var err error

// GormMysql 初始化Mysql数据库
func GormMysql() (db *gorm.DB) {
	conf := global.GVA_CONFIG.Mysql
	dbParams := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		conf.Username,
		conf.Password,
		conf.Path,
		conf.Port,
		conf.Dbname,
	)
	db, err = gorm.Open("mysql", dbParams)
	if err != nil {
		log.Fatal("mysql数据库连接失败：", err)
	}
	// 全局禁用表名复数
	db.SingularTable(true)

	db.DB().SetMaxIdleConns(conf.MaxIdleConns)
	db.DB().SetMaxOpenConns(conf.MaxOpenConns)
	fmt.Println("database init on port ", conf.Port)

	return
}
