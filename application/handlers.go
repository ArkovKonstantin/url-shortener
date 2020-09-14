package application

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"url-shortener/models"
)

func (app *Application) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app *Application) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(
			w,
			errors.Wrap(err, "reading request").Error(),
			http.StatusBadRequest,
		)
		return
	}
	defer r.Body.Close()
	var u models.URL
	err = json.Unmarshal(b, &u)
	if err != nil {
		http.Error(
			w,
			errors.Wrap(err, "unmarshal request data").Error(),
			http.StatusBadRequest,
		)
		return
	}
	shortURL, err := app.rep.Add(u.URL)
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	b, err = json.Marshal(ShortenResponse{shortURL})
	if err != nil {
		http.Error(
			w,
			errors.Wrap(err, "marshal response err").Error(),
			http.StatusInternalServerError,
		)
		return
	}
	w.WriteHeader(http.StatusCreated)
	w.Write(b)
}

func (app *Application) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	longURL, err := app.rep.GetURL(vars["short_url"])
	if err != nil {
		http.Error(
			w,
			err.Error(),
			http.StatusInternalServerError,
		)
		return
	}
	http.Redirect(w, r, longURL, http.StatusMovedPermanently)
}

type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}
