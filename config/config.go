package config

type Config struct {
	API     APIConfig
	MongoDB MongoDBConfig
	JWT     JWTConfig
}

type APIConfig struct {
	Port int
}

type MongoDBConfig struct {
	Host     string
	Port     uint16
	User     string
	Password string
}

type JWTConfig struct {
	Secret string
}
