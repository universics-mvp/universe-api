package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Env struct {
	AppEnv        string `mapstructure:"APP_ENV"`
	ServerAddress string `mapstructure:"SERVER_HOST"`
	Port          string `mapstructure:"PORT"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	ApiURL string `mapstructure:"API_URL"`

	YaGptOauthToken  string `mapstructure:"YA_OAUTH"`
	YaGptDirectoryID string `mapstructure:"YA_DIR_ID"`
}

func NewEnv() Env {
	env := Env{}

	viper.SetConfigFile(".env")

	_, err := os.Stat(".env")
	useEnvFile := !os.IsNotExist(err)

	if useEnvFile {
		viper.SetConfigType("env")
		viper.SetConfigName(".env")
		viper.AddConfigPath(".")

		err := viper.ReadInConfig()
		if err != nil {
			log.Fatal("Can't read the .env file: ", err)
		}
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&env)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if env.AppEnv != "production" {
		log.Println("The App is running in development env")
	}

	return env
}

var Module = fx.Options(
	fx.Provide(NewEnv),
)
