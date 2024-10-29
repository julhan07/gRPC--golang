package config

type Config struct {
	ServerPort        string
	ProductServiceURL string
	DBConnection      string
}

func LoadConfig() (*Config, error) {
	return &Config{
		ServerPort:        ":50052",
		ProductServiceURL: "localhost:50051",
		DBConnection:      "localhost:27017",
	}, nil
}
