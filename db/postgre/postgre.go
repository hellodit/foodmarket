package postgre

import (
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

//Connect to database
func Connect() *pg.DB {

	dbHost := viper.GetString("db_host")
	dbPort := viper.GetString("db_port")
	dbUser := viper.GetString("db_user")
	dbPass := viper.GetString("db_pass")
	dbName := viper.GetString("db_name")
	dbSslMode := viper.GetString("db_ssl_mode")

	parse := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", dbUser, dbPass, dbHost, dbPort, dbName, dbSslMode)
	opt, err := pg.ParseURL(parse)

	if err != nil {
		panic(err)
	}

	db := pg.Connect(opt)

	if db == nil {
		panic(fmt.Errorf("failed to connect database: %s", db))
	}

	logrus.Printf("Success connected to DB \n")

	return db
}
