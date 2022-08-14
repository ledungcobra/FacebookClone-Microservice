package config

import "os"

type config struct {
	MongoURL string
	SqlDsn   string
	Port     string
}

var Config *config

func Load() {

	defaultPort := os.Getenv("GATEWAY_PORT")
	if defaultPort == "" {
		defaultPort = "8080"
	}

	mongoUrl := os.Getenv("GATEWAY_MONGO_URL")
	if mongoUrl == "" {
		mongoUrl = "mongodb://localhost:27017"
	}

	sqlDsn := os.Getenv("GATEWAY_SQL_DSN")
	
	if sqlDsn == "" {
		sqlDsn = "root:root@tcp(localhost:3306)/gateway"
	}

	Config = &config{
		Port:     defaultPort,
		MongoURL: mongoUrl,
	}

}
