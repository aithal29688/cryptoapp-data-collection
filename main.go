package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/Crypto/CryptoDataCollection/misc"
	"github.com/Crypto/CryptoDataCollection/server"
)

var (
	showVersion = flag.Bool("version", false, "print version string")
	configFile  = flag.String("config", "dev.config.yaml", "Config file to use")
	hostname = "unknown"
)


func main() {
	flag.Parse()
	conf, err := misc.LoadConf(*configFile)
	if err != nil {
		log.WithField("error", err).Error("Failed to load config")
		os.Exit(-1)
	}

	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)

	if host, err := os.Hostname(); err == nil {
		hostname = host
	}

	p := &server.Loader{
		Conf: conf,
	}

	ticker := time.NewTicker(time.Minute)
	for {
		select {
			case <-ticker.C:
				p.HandleTick()

			case <-termChan:{
			shutdown(ticker, p)
			}
		}
	}

}

func shutdown(ticker *time.Ticker, p *server.Loader) {

	if ticker != nil {
		ticker.Stop()
	}

	os.Exit(0)
}


