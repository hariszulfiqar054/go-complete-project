package config

import "github.com/haris/go-rest/src/utils"

var (
	DB_NAME     = utils.GetEnv("DB_NAME", "go-rest")
	DB_PORT     = utils.GetEnv("DB_PORT", "3306")
	DB_USERNAME = utils.GetEnv("DB_USERNAME", "root")
	DB_PASSWORD = utils.GetEnv("DB_PASSWORD", "root")
	SERVER_PORT = utils.GetEnv("SERVER_PORT", "3000")
)
