package main

import (
	"log"
	"url-shortener/application"
	"url-shortener/models"
	"url-shortener/provider"
	"url-shortener/shortener"
)

var (
	config models.Config
)

func init() {
	err := models.LoadConfig(&config)
	if err != nil {
		log.Fatal(err)
	}
}
func main() {
	p := provider.New(&config.SQLDataBase)
	err := p.Open()

	if err != nil {
		log.Fatal(err)
	}

	rep := shortener.NewShortenerRepository(p, &config.Application)

	app := application.New(rep)
	app.Start()
}
