package main

import (
	"github.com/spf13/viper"
	"grahamcox.co.uk/worlds/service/internal"
	"grahamcox.co.uk/worlds/service/internal/service"
	"grahamcox.co.uk/worlds/service/internal/database"
)

func main() {
	config := viper.New()

	config.SetDefault("http.port", 3000)
	config.BindEnv("http.port", "PORT")

	config.BindEnv("pg.url", "PG_URI")

	configuration := internal.Config{
		HTTP: service.Config{
			Port: config.GetInt("http.port"),
		},
		Database: database.Config{
			URL: config.GetString("pg.url"),
		},
	}
	internal.Main(configuration)
}
