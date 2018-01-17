package server

import (
	"time"
	"io/ioutil"
	"net/http"
	"log"
	"fmt"
	"github.com/Crypto/cryptoapp-data-collection/misc"
	"github.com/Jeffail/gabs"
	"reflect"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

type Loader struct {
	LastUpdate time.Time
	Hostname   string
	Conf       *misc.Conf
}

func (p *Loader) HandleTick() {
	fromCurrencies := p.Conf.Currencies.FromCurrencyList
	toCurrencies := p.Conf.Currencies.ToCurrencyList
	url := constructUrl(fromCurrencies, toCurrencies)
	priceMulti := buildPriceObject(url)
	if priceMulti != nil {
		db, err := p.GetDbConn()
		if err != nil {
			log.Fatal(err)
		}

		defer db.Close()

		for _,fromCurrency := range strings.Split(fromCurrencies, ",") {
			for _,toCurrency := range strings.Split(toCurrencies, ","){
				insertStatement := buildInsertStatement()
				key := "RAW." + fromCurrency + "." + toCurrency + "."
				record := buildRecord(priceMulti, key)
				_, err := db.Exec(insertStatement,record.Type,record.Market,record.Fromsymbol,record.Tosymbol,record.Flags,record.Price,record.Lastupdate,record.Lastvolume,record.Lastvolumeto,record.Lasttradeid,record.Volumeday,record.Volumedayto,record.Volume24hour,record.Volume24hourto,record.Openday,record.Highday,record.Lowday,record.Open24hour,record.High24hour,record.Low24hour,record.Lastmarket,record.Change24hour,record.Changepct24hour,record.Changeday,record.Changepctday,record.Supply,record.Mktcap,record.Totalvolume24h,record.Totalvolume24hto)
				if err!=nil {
					panic(err)
				}

				log.Println("The prices for ", record.Fromsymbol," were last updated at :", time.Unix(record.Lastupdate,0))
			}
		}
	}

	p.LastUpdate = time.Now()
}

func buildRecord(priceMap map[string]interface{}, keyPrefix string) *misc.Record {
	return &misc.Record{
		Type: priceMap[keyPrefix+"TYPE"].(string),
		Market: priceMap[keyPrefix+"MARKET"].(string),
		Fromsymbol: priceMap[keyPrefix+"FROMSYMBOL"].(string),
		Tosymbol: priceMap[keyPrefix+"TOSYMBOL"].(string),
		Flags: priceMap[keyPrefix+"FLAGS"].(string),
		Price: priceMap[keyPrefix+"PRICE"].(float64),
		Lastupdate: int64(priceMap[keyPrefix+"LASTUPDATE"].(float64)),
		Lastvolume: priceMap[keyPrefix+"LASTVOLUME"].(float64),
		Lastvolumeto: priceMap[keyPrefix+"LASTVOLUMETO"].(float64),
		Lasttradeid: priceMap[keyPrefix+"LASTTRADEID"].(string),
		Volumeday: priceMap[keyPrefix+"VOLUMEDAY"].(float64),
		Volumedayto: priceMap[keyPrefix+"VOLUMEDAYTO"].(float64),
		Volume24hour: priceMap[keyPrefix+"VOLUME24HOUR"].(float64),
		Volume24hourto: priceMap[keyPrefix+"VOLUME24HOURTO"].(float64),
		Openday: priceMap[keyPrefix+"OPENDAY"].(float64),
		Highday: priceMap[keyPrefix+"HIGHDAY"].(float64),
		Lowday: priceMap[keyPrefix+"LOWDAY"].(float64),
		Open24hour: priceMap[keyPrefix+"OPEN24HOUR"].(float64),
		High24hour: priceMap[keyPrefix+"HIGH24HOUR"].(float64),
		Low24hour: priceMap[keyPrefix+"LOW24HOUR"].(float64),
		Lastmarket: priceMap[keyPrefix+"LASTMARKET"].(string),
		Change24hour: priceMap[keyPrefix+"CHANGE24HOUR"].(float64),
		Changepct24hour: priceMap[keyPrefix+"CHANGEPCT24HOUR"].(float64),
		Changeday: priceMap[keyPrefix+"CHANGEDAY"].(float64),
		Changepctday: priceMap[keyPrefix+"CHANGEPCTDAY"].(float64),
		Supply: priceMap[keyPrefix+"SUPPLY"].(float64),
		Mktcap: priceMap[keyPrefix+"MKTCAP"].(float64),
		Totalvolume24h: priceMap[keyPrefix+"TOTALVOLUME24H"].(float64),
		Totalvolume24hto: priceMap[keyPrefix+"TOTALVOLUME24HTO"].(float64),
	}
}

func buildInsertStatement() string {
	return `INSERT INTO cryptoprice (
	type,
	market,
	fromsymbol,
	tosymbol,
	flags,
	price,
	lastupdate,
	lastvolume,
	lastvolumeto,
	lasttradeid,
	volumeday,
	volumedayto,
	volume24hour,
	volume24hourto,
	openday,
	highday,
	lowday,
	open24hour,
	high24hour,
	low24hour,
	lastmarket,
	change24hour,
	changepct24hour,
	changeday,
	changepctday,
	supply,
	mktcap,
	totalvolume24h,
	totalvolume24hto) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17,$18,$19,$20,$21,$22,$23,$24,$25,$26,$27,$28,$29)
	RETURNING id`
}

func (p *Loader) GetDbConn() (*sql.DB, error){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		p.Conf.DBConn.Host, p.Conf.DBConn.Port, p.Conf.DBConn.User, p.Conf.DBConn.Password, p.Conf.DBConn.Dbname)
	return sql.Open("postgres", psqlInfo)
}

func constructUrl(fromCurr string, toCurr string) string {
	return "https://min-api.cryptocompare.com/data/pricemultifull?fsyms=" + fromCurr + "&tsyms=" + toCurr
}

func buildPriceObject(url string) map[string]interface{} {
	httpClient := http.Client{
		Timeout: time.Second * 10, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		log.Println(getErr)
		log.Println("Skipping the records for time : ", time.Now())
		return nil
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	m, err := FlattenJson(body)
	if err != nil {
		//TODO Log error
	}

	return m

}

func FlattenJson(data []byte) (map[string]interface{}, error) {
	jParsed, err := gabs.ParseJSON(data)
	if err != nil {
		//TODO: log error
		return nil, err
	}

	jp, err := jParsed.ChildrenMap()
	if err != nil {
		//TODO: log error
		return nil, err
	}

	return BuildJsonMap(jp), nil
}

func BuildJsonMap(outterChild map[string] *gabs.Container) map[string]interface{} {
	var output map[string]interface{}
	output = make(map[string]interface{})

	for key, value := range outterChild {
		if reflect.ValueOf(value.Data()).Kind() == reflect.Map {
			if oChild, err := value.ChildrenMap(); err != nil {
				fmt.Println(err)
			} else {
				rOutput := BuildJsonMap(oChild)
				for k, v := range rOutput {
					if(!strings.HasPrefix(key, "DISPLAY")){
						output[key+"."+k] = v
					}
				}
			}
		} else {
			output[key] = value.Data()
		}
	}

	return output
}