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
	models.LoadConfig(&config)
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
