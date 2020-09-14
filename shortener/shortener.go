package shortener

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"url-shortener/models"
	"url-shortener/provider"
)

type Shortener interface {
	Add(string) (string, error)
	GetURL(string) (string, error)
}

type shortener struct {
	p      provider.Provider
	appCfg *models.Application
}

func NewShortenerRepository(p provider.Provider, appCfg *models.Application) Shortener {
	return &shortener{p, appCfg}
}

func (s *shortener) Add(longURL string) (string, error) {
	db, err := s.p.GetConn()
	if err != nil {
		return "", errors.Wrap(err, "get db connection err:")
	}
	tx, err := db.Begin()
	if err != nil {
		return "", errors.Wrap(err, "begin tx err:")
	}
	var id int
	var shortURL string
	{
		stmt, err := tx.Prepare("INSERT INTO url (long_url) VALUES ($1) RETURNING id")
		if err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "tx prepare err:")
		}
		defer stmt.Close()
		err = stmt.QueryRow(longURL).Scan(&id)
		if err != nil {
			return "", errors.Wrap(err, "insert operation err:")
		}
	}
	{
		stmt, err := tx.Prepare("UPDATE  url SET short_url = ($1) where id = ($2)")
		if err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "tx prepare err:")
		}
		defer stmt.Close()
		num, err := convertBase(strconv.Itoa(id), 10, 66)
		if err != nil {
			return "", errors.Wrap(err, "convert id to base 66 num err:")
		}
		shortURL = fmt.Sprintf("http://%s:%d/%s", s.appCfg.Host, s.appCfg.Port, num)
		if _, err := stmt.Exec(num, id); err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "update operation err:")
		}
	}
	err = tx.Commit()
	if err != nil {
		return "", errors.Wrap(err, "commit err:")
	}

	return shortURL, nil
}

func (s *shortener) GetURL(path string) (string, error) {
	db, err := s.p.GetConn()
	if err != nil {
		return "", errors.Wrap(err, "get db connection err:")
	}
	row := db.QueryRow("SELECT long_url FROM url WHERE short_url = ($1)", path)
	var longURL string
	err = row.Scan(&longURL)
	if err == sql.ErrNoRows {
		return "", errors.Wrap(err, "Not Found")
	} else if err != nil {
		return "", errors.Wrap(err, "select operation err:")
	}
	return longURL, nil
}
