package shortener

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"url-shortener/models"
	"url-shortener/provider"
)

type Shortener interface {
	AddURL(string) (string, error)
	AddURLWithName(string, string) error
	GetURLByName(string) (string, error)
}

type shortener struct {
	p      provider.Provider
	appCfg *models.Application
}

func NewShortenerRepository(p provider.Provider, appCfg *models.Application) Shortener {
	return &shortener{p, appCfg}
}

func (s *shortener) AddURL(url string) (string, error) {
	db, err := s.p.GetConn()
	if err != nil {
		return "", errors.Wrap(err, "get db connection err:")
	}

	tx, err := db.Begin()
	if err != nil {
		return "", errors.Wrap(err, "begin tx err:")
	}

	var id int
	var name, shortURL string
	{
		stmt, err := tx.Prepare("INSERT INTO url (url) VALUES ($1) RETURNING id")
		if err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "tx prepare err:")
		}
		defer stmt.Close()
		err = stmt.QueryRow(url).Scan(&id)
		if err != nil {
			return "", errors.Wrap(err, "insert operation err:")
		}
	}
	{
		stmt, err := tx.Prepare("UPDATE url SET name = ($1) where id = ($2)")
		if err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "tx prepare err:")
		}
		defer stmt.Close()
		name, err = convertBase(strconv.Itoa(id), 10, 66)
		if err != nil {
			return "", errors.Wrap(err, "convert id to short_url err:")
		}
		if _, err := stmt.Exec(name, id); err != nil {
			tx.Rollback()
			return "", errors.Wrap(err, "update operation err:")
		}
	}
	err = tx.Commit()
	if err != nil {
		return "", errors.Wrap(err, "commit err:")
	}

	shortURL = fmt.Sprintf("http://%s:%d/%s", s.appCfg.Host, s.appCfg.Port, name)

	return shortURL, nil
}

func (s *shortener) AddURLWithName(url, name string) error {
	db, err := s.p.GetConn()
	if err != nil {
		return errors.Wrap(err, "get db connection err:")
	}

	result, err := db.Exec("INSERT INTO url (url, name) VALUES($1, $2)", url, name)
	if err != nil {
		return errors.Wrap(err, "insert operation err:")
	}

	_, err = result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "rowsAffected operation err:")
	}

	return nil
}

func (s *shortener) GetURLByName(name string) (string, error) {
	db, err := s.p.GetConn()
	if err != nil {
		return "", errors.Wrap(err, "get db connection err:")
	}

	var url string
	row := db.QueryRow("SELECT url FROM url WHERE name = ($1)", name)
	err = row.Scan(&url)
	if err != nil {
		return "", err
	}

	return url, nil
}
