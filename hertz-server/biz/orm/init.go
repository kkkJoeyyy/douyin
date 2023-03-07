package orm

import (
	"fmt"
	"github.com/simplecolding/douyin/hertz-server/biz/orm/dal"
	"github.com/simplecolding/douyin/hertz-server/biz/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init()  {
	db, err := gorm.Open(mysql.Open(utils.MySQLDefaultDSN))
	if err != nil {
		panic(fmt.Errorf("连接mysql失败：%w", err))
	}
	dal.SetDefault(db)
}