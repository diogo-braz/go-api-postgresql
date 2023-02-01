package configs

import "github.com/spf13/viper"

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Hostname string
	Port     string
	User     string
	Password string
	Database string
}

func init() {
	viper.SetDefault("api.port", "9000")
	viper.SetDefault("database.hostname", "localhost")
	viper.SetDefault("database.port", "5432")
}

func Load() error {
	viper.SetConfigFile("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config)

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Hostname: viper.GetString("database.Hostname"),
		Port:     viper.GetString("database.Port"),
		User:     viper.GetString("database.User"),
		Password: viper.GetString("database.Password"),
		Database: viper.GetString("database.Database"),
	}

	return nil
}

func GetServerPort() string {
	return cfg.API.Port
}

func GetDB() DBConfig {
	return cfg.DB
}
