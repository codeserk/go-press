package config

type Config struct {
	Api     ApiConfig
	MongoDB MongoDBConfig
	JWT 	JWTConfig
}

type ApiConfig struct {
	Port int
}

type MongoDBConfig struct {
	Host     string
	User     string
	Password string
}

type JWTConfig struct {
	Secret string
}