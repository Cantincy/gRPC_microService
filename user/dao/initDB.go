package dao

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strings"
)

var DB *gorm.DB

func InitDB() error {
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.port")
	database := viper.GetString("mysql.database")
	username := viper.GetString("mysql.username")
	password := viper.GetString("mysql.password")
	charset := viper.GetString("mysql.charset")
	dsn := strings.Join([]string{username, ":", password, "@(", host, ":", port, ")/", database, "?charset=", charset, "&parseTime=True&loc=Local"}, "")
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return err
}
