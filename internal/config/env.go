package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/fx"
)

type Env struct {
	AppEnv string `mapstructure:"APP_ENV"`
	Port   string `mapstructure:"PORT"`

	DBHost string `mapstructure:"DB_HOST"`
	DBPort string `mapstructure:"DB_PORT"`
	DBUser string `mapstructure:"DB_USER"`
	DBPass string `mapstructure:"DB_PASS"`
	DBName string `mapstructure:"DB_NAME"`

	ApiURL string `mapstructure:"API_URL"`

	BotToken         string `mapstructure:"BOT_TOKEN"`
	YaGptOauthToken  string `mapstructure:"YA_OAUTH"`
	YaGptDirectoryID string `mapstructure:"YA_DIR_ID"`
}

func NewEnv() Env {
	env := Env{}

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

		err = viper.Unmarshal(&env)
		if err != nil {
			log.Fatal("Environment can't be loaded: ", err)
		}
	} else {
		env.bindEnv()
	}

	if env.AppEnv != "production" {
		log.Println("The App is running in development env")
	}

	return env
}

func (e *Env) bindEnv() {
	e.ApiURL = os.Getenv("API_URL")
	e.AppEnv = os.Getenv("APP_ENV")
	e.Port = os.Getenv("PORT")

	e.DBHost = os.Getenv("DB_HOST")
	e.DBPort = os.Getenv("DB_PORT")
	e.DBUser = os.Getenv("DB_USER")
	e.DBPass = os.Getenv("DB_PASS")
	e.DBName = os.Getenv("DB_NAME")

	e.YaGptOauthToken = os.Getenv("YA_OAUTH")
	e.YaGptDirectoryID = os.Getenv("YA_DIR_ID")
	e.BotToken = os.Getenv("BOT_TOKEN")
}

var Module = fx.Options(
	fx.Provide(NewEnv),
)
