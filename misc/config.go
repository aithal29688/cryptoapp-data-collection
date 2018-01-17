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
	DBConn		DBConn
}
type Currencies struct {
	FromCurrencyList	string
	ToCurrencyList		string
}

type DBConn struct {
	Host		string
	Password	string
	Port 		int
	User 		string
	Dbname 		string
}

type Record struct {
	Type     string
	Market     string
	Fromsymbol     string
	Tosymbol     string
	Flags     string
	Price     float64
	Lastupdate     int64
	Lastvolume     float64
	Lastvolumeto     float64
	Lasttradeid     string
	Volumeday     float64
	Volumedayto     float64
	Volume24hour     float64
	Volume24hourto     float64
	Openday     float64
	Highday     float64
	Lowday     float64
	Open24hour     float64
	High24hour     float64
	Low24hour     float64
	Lastmarket     string
	Change24hour     float64
	Changepct24hour     float64
	Changeday     float64
	Changepctday     float64
	Supply     float64
	Mktcap     float64
	Totalvolume24h     float64
	Totalvolume24hto     float64
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