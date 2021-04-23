package config

type Config struct {
	Url              string
	DynamicsUserName string
	DynamicsPassword string
	Mode             string
	Host             string
	Port             string
}


func LoadConfig() *Config {
	return &Config{}
}