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
	User     string
	Password string
}

type JWTConfig struct {
	Secret string
}
