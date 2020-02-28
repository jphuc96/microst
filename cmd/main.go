package main

import (
	"os"

	"github.com/jphuc96/microst/pkg/db/mongodb"
	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/jphuc96/microst/pkg/api"
)

func main() {
	// app.InitConfig()
	// viper.SetConfigFile("/app/microst.yaml")

	// uri := viper.GetString("mongodb.uri")
	uri := os.Getenv("MONGODB_URI")

	_, err := mongodb.New(uri)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("Successully init dbs")

	e := echo.New()
	api.Init(e)
	api.Start(e)
}
