package sql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Name     string
}

func NewConn() (db *sql.DB, cf func(), err error) {
	dbViper := viper.New()
	dbViper.AddConfigPath(".")

	dbViper.SetConfigFile("db.yaml")
	if err = dbViper.ReadInConfig(); err != nil {
		return
	}
	var config Config
	if err = dbViper.UnmarshalKey("db", &config); err != nil {
		return
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", config.Username, config.Password, config.Host, config.Port, config.Name)
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return nil, nil, err
	}
	cf = func() { _ = db.Close() }
	return
}
