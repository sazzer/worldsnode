package main

import (
	"github.com/spf13/viper"
	"grahamcox.co.uk/worlds/service/internal"
	"grahamcox.co.uk/worlds/service/internal/service"
)

func main() {
	config := viper.New()

	config.SetDefault("http.port", 3000)
	config.BindEnv("http.port", "PORT")

	configuration := internal.Config{
		HTTP: service.Config{
			Port: config.GetInt("http.port"),
		},
	}
	internal.Main(configuration)
}
