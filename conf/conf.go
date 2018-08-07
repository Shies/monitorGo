package conf

import (
	"flag"
	"os"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

var (
	Conf Config
)

type Config struct {
	Db  *Db  `toml:"db"`
	Log *Log `toml:"xlog"`
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

func Logger(dir string) (err error) {
	flag.Parse()
	outfile, err := os.OpenFile(dir, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(*outfile, "open failed")
		os.Exit(1)
	}
	log.SetOutput(outfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	return
}

