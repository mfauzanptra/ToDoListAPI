package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBUser string
	DBPass string
	DBHost string
	DBPort int
	DBName string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}
	isRead := true

	if val, found := os.LookupEnv("MYSQL_USER"); found {
		app.DBUser = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PASSWORD"); found {
		app.DBPass = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_HOST"); found {
		app.DBHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.DBPort = cnv
		isRead = false
	}
	if val, found := os.LookupEnv("MYSQL_DBNAME"); found {
		app.DBName = val
		isRead = false
	}

	if isRead {
		err := godotenv.Load("local.env")
		if err != nil {
			fmt.Println("Error saat baca env", err.Error())
			return nil
		}

		app.DBUser = os.Getenv("MYSQL_USER")
		app.DBPass = os.Getenv("MYSQL_PASSWORD")
		app.DBHost = os.Getenv("MYSQL_HOST")
		readData := os.Getenv("MYSQL_PORT")
		app.DBPort, err = strconv.Atoi(readData)
		if err != nil {
			fmt.Println("Error saat convert", err.Error())
			return nil
		}
		app.DBName = os.Getenv("MYSQL_DBNAME")
	}

	return &app
}
