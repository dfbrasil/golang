package config

// Properties Configuration properties based on environment variables
type Properties struct {
	Port string `env:"PORT" env-default:"8080"`
	Host string `env:"HOST" env-default:"localhost"`
	DBHost string `env:"DB_HOST" env-default:"localhost"`
	DBPort string `env:"DB_PORT" env-default:"27017"`
	DBName string `env:"DB_NAME" env-default:"tronicscorp"`
	ProductCollection string `env:"PRODUCTS_COL_NAME" env-default:"products"`
	UsersCollection string `env:"USERS_COL_NAME" env-default:"users"`
	JwtTokenSecret string `env:"JWT_TOKEN_SECRET" env-default:"senhasecreta"`
}