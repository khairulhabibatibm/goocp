package config

import (
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
}

func BuildDBConfig() *DBConfig {
	hostmysql, _ := os.LookupEnv("MYSQL_HOST")
	usermysql, _ := os.LookupEnv("MYSQL_USER")
	passmysql, _ := os.LookupEnv("MYSQL_PASS")
	dbmysql, _ := os.LookupEnv("MYSQL_DB")
	dbConfig := DBConfig{
		Host:     hostmysql,
		Port:     3306,
		User:     usermysql,
		DBName:   dbmysql,
		Password: passmysql,
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.DBName,
	)
}
