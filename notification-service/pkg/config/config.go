package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoURL    string
	SqlDsn      string
	ServerPort  string
	GatewayCost int
	Env         string
	JWTSecret   string
}

var Cfg *Config

func Init() {
	Cfg = &Config{}
	log.Println("Loading env from .env file")
	err := godotenv.Load(".env")
	if err != nil {
		log.Print("Error loading .env file", err)
	}
	Cfg.loadServerConfig()
	Cfg.loadMongoDbURL()
	Cfg.loadPostgresURL()
	Cfg.GatewayCost = Cfg.mustLoadEnvInt("GATEWAY_BCRYPT_COST")
}

func (c *Config) loadPostgresURL() {
	pgUser := c.mustLoadEnv("GATEWAY_POSTRES_USER")
	pgPassword := c.mustLoadEnv("GATEWAY_POSTRES_PASSWORD")
	pgDatabase := c.mustLoadEnv("GATEWAY_POSTRES_DATABASE")
	pgHost := c.mustLoadEnv("GATEWAY_POSTRES_HOST")
	pgPort := c.mustLoadEnv("GATEWAY_POSTRES_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", pgHost, pgUser, pgPassword, pgDatabase, pgPort)
	c.SqlDsn = dsn
}

func (c *Config) loadMongoDbURL() {
	c.MongoURL = c.loadEnvOrDefault("GATEWAY_MONGO_URL", "mongodb://localhost:27017")
}

func (c *Config) loadServerConfig() {
	c.ServerPort = c.loadEnvOrDefault("GATEWAY_PORT", "8080")
	c.Env = c.loadEnvOrDefault("GATEWAY_ENV", "development")
	c.JWTSecret = c.mustLoadEnv("GATEWAY_JWT_SECRET")
}

func (c *Config) loadEnvOrDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func (c *Config) mustLoadEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}
	return value
}

func (c *Config) mustLoadEnvInt(key string) int {
	value := os.Getenv(key)
	if value == "" {
		log.Fatal("Environment variable " + key + " is not set")
	}
	parsedInt, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("Environment variable " + key + " is not a valid integer")
	}
	return parsedInt
}
