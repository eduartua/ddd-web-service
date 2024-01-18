package app

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type PostgresConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func (c PostgresConfig) ConnectionInfo() string {
	if c.Password == "" {
		return fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Name)
	}
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", c.Host, c.Port, c.User, c.Password, c.Name)
}

func DefaultPostgresConfig() PostgresConfig {
	return PostgresConfig{
		Host:     "localhost",
		Port:     5432,
		User:     "",
		Password: "",
		Name:     "",
	}
}

func LoadConfig(configReq bool) Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		if configReq {
			panic(err)
		}
		fmt.Println("Using the default config...")
		return DefaultConfig()
	}

	var c Config
	dec := yaml.NewDecoder(f)
	err = dec.Decode(&c)
	if err != nil {
		panic(err)
	}
	return c
}

type Config struct {
	Database PostgresConfig `yaml:"database"`
	Vars     Vars           `yaml:"vars"`
}

func (c Config) IsProd() bool {
	return c.Vars.Env == "prod"
}

func DefaultConfig() Config {
	return Config{
		Database: DefaultPostgresConfig(),
	}
}

type Vars struct {
	Env              string `yaml:"env"`
	HMACKey          string `yaml:"hmac_key"`
	Pepper           string `yaml:"pepper"`
}