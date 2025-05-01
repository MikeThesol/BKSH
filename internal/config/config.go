package config

type Config struct {
	Env        string
	ServerPort string
	DB
}

type DB struct {
	DBName   string
	User     string
	Host     string
	Password string
	SSLMode  string
	Port     string
}
