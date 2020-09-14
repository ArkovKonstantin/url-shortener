package models

import (
	"github.com/BurntSushi/toml"
	"io/ioutil"
	"strings"
	"time"
)

var (
	configPath = "config/config.toml"
	hashPaths  = []string{configPath}
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
func LoadConfig(c *Config) {
	_, err := toml.DecodeFile(configPath, c)
	if err != nil {
		return
	}
	// c.SQLDataBase.User = getCredential("/etc/scrt/chat-server/sqlUser")
	// c.SQLDataBase.Password = getCredential("/etc/scrt/chat-server/sqlPassword")

}

func getCredential(path string) string {
	hashPaths = append(hashPaths, path)
	c, _ := ioutil.ReadFile(path)
	return strings.TrimSpace(string(c))
}
