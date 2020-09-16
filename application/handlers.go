package application

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	errorURLNotFound = "url does not exists"
	errorDuplicateName = "url with this name already exists"
)

func (app *Application) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (app *Application) ShortenHandler(w http.ResponseWriter, r *http.Request) {
	req, err := decodeShortenReq(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	u, err := app.rep.AddURL(req.URL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := encodeResponse(w, ShortenResponse{u}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *Application) CreateHandler(w http.ResponseWriter, r *http.Request) {
	req, err := decodeCreateReq(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_ , err = app.rep.GetURLByName(req.Name)
	if err == sql.ErrNoRows {
		err := app.rep.AddURLWithName(req.URL, req.Name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Error(w, errorDuplicateName, http.StatusBadRequest)
}

func (app *Application) RedirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	url, err := app.rep.GetURLByName(vars["name"])

	if err == sql.ErrNoRows {
		http.Error(w, errorURLNotFound, http.StatusNotFound)
		return
	} else if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusMovedPermanently)
}
