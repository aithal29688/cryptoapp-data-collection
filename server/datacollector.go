package server

import (
	"time"
	"github.com/Crypto/CryptoDataCollection/misc"
	"github.com/Crypto/CryptoDataCollection/models"
	"io/ioutil"
	"net/http"
	"log"
	"encoding/json"
	"fmt"
)

type Loader struct {
	LastUpdate time.Time
	Hostname   string
	Conf       *misc.Conf
}

func (p *Loader) HandleTick() {
	url := constructUrl(p.Conf.Currencies.Currencylist)
	priceMulti := buildPriceObject(url)

	fmt.Println("the bitcoin price is : " , priceMulti.Raw.BTC.USD.Price)

	p.LastUpdate = time.Now()
}

func constructUrl(fromCurr string) string {
	return "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=" + fromCurr + "&tsyms=USD"
}

func buildPriceObject(url string) models.PriceMulti {
	httpClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	priceMulti := models.PriceMulti{}
	jsonErr := json.Unmarshal(body, &priceMulti)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	return priceMulti

}