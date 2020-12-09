package postgre

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

//Connect to database
func Connect() *pg.DB {

	dbHost := viper.GetString("db_host")
	dbPort := viper.GetString("db_port")
	dbUser := viper.GetString("db_user")
	dbPass := viper.GetString("db_pass")
	dbName := viper.GetString("db_name")
	dbSslMode := viper.GetString("db_sslmode")

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
