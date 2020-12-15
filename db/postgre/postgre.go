package postgre

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//Connect to database
func Connect() *gorm.DB {

	dbHost := viper.GetString("db_host")
	dbPort := viper.GetString("db_port")
	dbUser := viper.GetString("db_user")
	dbPass := viper.GetString("db_pass")
	dbName := viper.GetString("db_name")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	fmt.Print(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
