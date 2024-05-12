package config

import (
	"fmt"
	"log"
	"os"

	"github.com/backend-magang/eniqilo-store/utils/pkg"
	"github.com/joho/godotenv"
)

type Config struct {
	AppHost    string `mapstructure:"APP_HOST"`
	AppPort    string `mapstructure:"APP_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUsername string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBParams   string `mapstructure:"DB_PARAMS"`
	DBSchema   string `mapstructure:"DB_SCHEMA"`
	JWTSecret  string `mapstructure:"JWT_SECRET"`
	BCryptSalt string `mapstructure:"BCRYPT_SALT"`
	SqlTrx     *pkg.SqlWithTransactionService
}

func Load() (conf Config) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, err: ", err.Error())
	}

	conf = Config{
		AppHost:    os.Getenv("APP_HOST"),
		AppPort:    os.Getenv("APP_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBPort:     os.Getenv("DB_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBParams:   os.Getenv("DB_PARAMS"),
		DBSchema:   "public",
		JWTSecret:  os.Getenv("JWT_SECRET"),
		BCryptSalt: os.Getenv("BCRYPT_SALT"),
	}

	return
}

func (cfg *Config) GetDSN() (dsn string) {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?%s&search_path=%s",
		cfg.DBUsername,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
		cfg.DBParams,
		cfg.DBSchema,
	)
}
