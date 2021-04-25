package config

type Config struct {
	Api     ApiConfig
	MongoDB MongoDBConfig
}

type ApiConfig struct {
	Port int
}

type MongoDBConfig struct {
	Host     string
	User     string
	Password string
}
