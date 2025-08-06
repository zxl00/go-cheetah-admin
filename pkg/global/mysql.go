/*
@auth: AnRuo
@source: 蜀道运维 公众号
@time: 2025/8/6
*/

package global

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	modelSystem "github.com/zxl00/go-cheetah-admin/internal/model/system"
)

var (
	GORM *gorm.DB
	err  error
)

// 初始化数据库

func InitMysql() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.DbUser"),
		viper.GetString("mysql.DbPwd"),
		viper.GetString("mysql.DbHost"),
		viper.GetInt("mysql.DbPort"),
		viper.GetString("mysql.DbName"))
	GORM, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic("数据库连接失败" + err.Error())
	}
	if viper.GetInt("mysql.ActiveDebug") == 1 {
		GORM = GORM.Debug()
	}
	// 开启连接池
	db, _ := GORM.DB()
	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	if err = db.Ping(); err != nil {
		panic("数据库连接失败" + err.Error())

	}
	_ = GORM.AutoMigrate(
		modelSystem.User{},
		modelSystem.Role{},
		modelSystem.Menu{},
		modelSystem.APIs{},
		modelSystem.Ldap{},
	)
}
