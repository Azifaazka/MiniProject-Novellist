package config

type Config struct {
	DB_USERNAME,DB_PASSWORD,DB_PORT,DB_HOST,DB_NAME string
}

func InitConfiguration() Config {
	return Config{
		DB_USERNAME: "root",
		DB_PASSWORD: "123456",
		DB_NAME: "novellist",
		DB_PORT: "3306",
		DB_HOST: "0.0.0.0",
	}
}