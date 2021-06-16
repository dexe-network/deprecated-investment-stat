package api

type CORSConfig struct {
	AllowCredentials bool     `env:"CORS_ALLOW_CREDENTIALS" envDefault:"true"`
	AllowedHeaders   []string `env:"CORS_ALLOWED_HEADERS" envSeparator:"," envDefault:"Origin,X-Requested-With,Content-Type,Accept,Authorization,User-Agent,X-Refresh-Token"`
	ExposedHeaders   []string `env:"CORS_EXPOSED_HEADERS" envSeparator:"," envDefault:"Origin,X-Requested-With,Content-Type,Accept,User-Agent"`
	AllowedMethods   []string `env:"CORS_ALLOWED_METHODS" envSeparator:"," envDefault:"GET,POST,PUT,PATCH,DELETE,OPTIONS,HEAD"`
	AllowedOrigins   []string `env:"CORS_ALLOWED_ORIGINS" envSeparator:"," envDefault:"http://127.0.0.1:3000,http://localhost:3000,https://localhost:3000"`
}
