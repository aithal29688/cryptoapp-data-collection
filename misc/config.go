package misc

import (
	"io/ioutil"
	"gopkg.in/yaml.v1"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

type Log struct {
	Level       string
	Filename    string
	MaxSizeMB   int
	MaxBackups  int
	MaxAgeDays  int
	WriteStdout bool
	Json        bool
}
type Http struct {
	Address      string
	ReadTimeout  int
	WriteTimeout int
}

type Conf struct {
	Log             Log
	Http            Http
	Environment     string
	RefreshInterval int
	StatsdHost      string
	Currencies 	Currencies
}
type Currencies struct {
	Currencylist	string
}

func LoadConf(filename string) (*Conf, error) {
	cnf := Conf{}

	source, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(source, &cnf); err != nil {
		return nil, err
	}

	log.Info("Loaded config file")

	return &cnf, nil
}