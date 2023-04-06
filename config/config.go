package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

var (
	JWT_SECRET string
)

type AppConfig struct {
	DBUser     string
	DBPassword string
	DBHost     string
	DBPort     uint16
	DBName     string
}

func InitConfig() *AppConfig {
	var cnf = readConfig()
	if cnf == nil {
		return nil
	}

	return cnf
}

func readConfig() *AppConfig {
	var result = new(AppConfig)

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Cannot read config variable")
		return nil
	}

	result.DBUser = os.Getenv("DBUser")
	result.DBPassword = os.Getenv("DBPassword")
	result.DBHost = os.Getenv("DBHost")
	cnvPort, err := strconv.Atoi(os.Getenv("DBPort"))
	if err != nil {
		log.Error("Cannot convert database port", err.Error())
		return nil
	}
	result.DBPort = uint16(cnvPort)
	result.DBName = os.Getenv("DBName")
	JWT_SECRET = os.Getenv("JWT_SECRET")
	return result
}
