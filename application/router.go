package application


func (app *Application) router() {
	app.r.HandleFunc("/health", app.HealthHandler)
	app.r.HandleFunc("/shorten", app.ShortenHandler)
	app.r.HandleFunc("/{short_url}", app.RedirectHandler)

}
