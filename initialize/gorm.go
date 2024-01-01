package initialize

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"wechat/global"
)

// Gorm 初始化数据库并产生数据库全局变量
func Gorm() *gorm.DB {
	switch global.GVA_CONFIG.SYSTEM.DbType {
	case "mysql":
		return GormMysql()
	default:
		return GormMysql()
	}
}
