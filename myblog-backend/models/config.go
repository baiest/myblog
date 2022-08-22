package models

type DatabaseConfig struct {
	User     string
	Password string
	Port     string
	DbName   string
}

type ConfigServer struct {
	Port           string
	JwtKey         string
	DatabaseConfig DatabaseConfig
}
