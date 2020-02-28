package main

import (
	"github.com/jphuc96/microst/pkg/db/mongodb"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/jphuc96/microst/pkg/api"
)

func main() {
	// app.InitConfig()
	viper.SetConfigFile("/app/microst.yaml")

	uri := viper.GetString("mongodb.uri")

	_, err := mongodb.New(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Successully init dbs")

	e := echo.New()
	api.Init(e)
	api.Start(e)
}
