package conf

import (
	"github.com/BurntSushi/toml"
)

var (
	Conf Config
)

type Config struct {
	Db  *Db  `toml:"db"`
	Log *Log `toml:"log"`
}

// Log
type Log struct {
	Dir	string	`toml:"dir"`
}

// DB
type Db struct {
	Addr     string `toml:"addr"`
	Port     string `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Db_name  string `toml:"db_name"`
}

func ParseConfig() (err error){
	 _, err = toml.DecodeFile("./config.example.toml", &Conf)
	return
}
