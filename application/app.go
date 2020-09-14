package application

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"url-shortener/shortener"
)

type Application struct {
	servicePort int
	r           *mux.Router
	rep         shortener.Shortener
}

func New(rep shortener.Shortener) Application {
	return Application{servicePort: 9000, r: mux.NewRouter(), rep: rep}
}

func (app *Application) Start() {
	app.router()
	fmt.Println("start and listening on :9000 ...")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(`:%d`, app.servicePort), app.r))
}

