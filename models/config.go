package models

import (
	"github.com/BurntSushi/toml"
	"os"
	"time"
)

var (
	devConfigPath  = "config/config.dev.toml"
	prodConfigPath = "config/config.prod.toml"
)

type duration time.Duration

// SQLDataBase struct
type SQLDataBase struct {
	Server          string   `toml:"Server"`
	Database        string   `toml:"Database"`
	ApplicationName string   `toml:"ApplicationName"`
	MaxIdleConns    int      `toml:"MaxIdleConns"`
	MaxOpenConns    int      `toml:"MaxOpenConns"`
	ConnMaxLifetime duration `toml:"ConnMaxLifetime"`
	Port            int      `toml:"Port"`
	User            string   `toml:"User"`
	Password        string   `toml:"Password"`
}

type Application struct {
	Host string `toml:"Host"`
	Port int    `toml:"Port"`
}

// Config struct
type Config struct {
	Application Application `toml:"Application"`
	SQLDataBase SQLDataBase `toml:"SQLDataBase"`
	ServerOpt   ServerOpt   `toml:"ServerOpt"`
	HashSum     []byte
}

func (d *duration) UnmarshalText(text []byte) error {
	temp, err := time.ParseDuration(string(text))
	*d = duration(temp)
	return err
}

// ServerOpt struct
type ServerOpt struct {
	ReadTimeout  duration
	WriteTimeout duration
	IdleTimeout  duration
}

// LoadConfig from path
func LoadConfig(c *Config) error {
	env := os.Getenv("ENV")
	var p string
	if env == "" || env == "dev" {
		p = devConfigPath
	} else if env == "prod" {
		p = prodConfigPath
	}
	_, err := toml.DecodeFile(p, c)
	if err != nil {
		return err
	}
	return nil
}
