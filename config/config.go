package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ENV                 string
	SERVER_PORT         string
	USER_TABLE          string
	LOGGER_URL          string
	SECRET_KEY          string
	POSTGRES_DB         string
	POSTGRES_USER       string
	POSTGRES_HOST       string
	POSTGRES_PORT       string
	POSTGRES_PASSWORD   string
	ARTICLE_SERVICE_URL string
	DEBUG               bool
	TEST                bool
}

func NewConfig() (*Config, error) {
	ENV := os.Getenv("ENV")
	switch ENV {
	case "development":
		err := godotenv.Load(".env")
		if err != nil {
			return nil, err
		}
	}

	var (
		SECRET_KEY          = os.Getenv("SECRET_KEY")
		POSTGRES_PASSWORD   = os.Getenv("POSTGRES_PASSWORD")
		POSTGRES_USER       = "postgres"
		POSTGRES_DB         = "postgres"
		POSTGRES_HOST       = "postgres"
		POSTGRES_PORT       = "5432"
		SERVER_PORT         = "8001"
		USER_TABLE          = "Users"
		LOGGER_URL          = "http://logger:8002/v1/logger"
		ARTICLE_SERVICE_URL = "http://articles:8001/v1/articles"
		DEBUG               = false
		TEST                = false
	)

	switch ENV {
	case "production":
		TEST = false
		DEBUG = false

	case "production_test":
		TEST = true
		DEBUG = true
		USER_TABLE = "PriductionTestUsers"

	case "development":
		TEST = true
		DEBUG = true
		POSTGRES_HOST = "localhost"
		USER_TABLE = "DevUsers"
		LOGGER_URL = "http://localhost:8002/v1/logger"
		ARTICLE_SERVICE_URL = "http://localhost:8001/v1/articles"

	case "development_test":
		TEST = true
		DEBUG = true
		SECRET_KEY = "testsecret"
		POSTGRES_PASSWORD = "pass1234"
		POSTGRES_HOST = "localhost"
		USER_TABLE = "TestUsers"
		LOGGER_URL = "http://localhost:8002/v1/logger"
		ARTICLE_SERVICE_URL = "http://localhost:8001/v1/articles"

	case "docker":
		TEST = true
		DEBUG = true
		USER_TABLE = "DockerUsers"
		LOGGER_URL = "http://logger:8002/v1/logger"

	case "docker_test":
		TEST = true
		DEBUG = true
		USER_TABLE = "DockerUsers"
		LOGGER_URL = "http://logger:8002/v1/logger"
	}

	config := Config{
		ENV:                 ENV,
		SERVER_PORT:         SERVER_PORT,
		USER_TABLE:          USER_TABLE,
		SECRET_KEY:          SECRET_KEY,
		LOGGER_URL:          LOGGER_URL,
		DEBUG:               DEBUG,
		TEST:                TEST,
		POSTGRES_DB:         POSTGRES_DB,
		POSTGRES_USER:       POSTGRES_USER,
		POSTGRES_HOST:       POSTGRES_HOST,
		POSTGRES_PORT:       POSTGRES_PORT,
		POSTGRES_PASSWORD:   POSTGRES_PASSWORD,
		ARTICLE_SERVICE_URL: ARTICLE_SERVICE_URL,
	}

	return &config, nil
}
