package models


type PriceMulti struct {
	Raw	Raw	`json:"RAW"`
}

type Raw struct {
	ETH	ETH	`json:"ETH"`
	BTC	BTC	`json:"BTC"`
	LTC	LTC	`json:"LTC"`
	BCH	BCH	`json:"BCH"`
}

type USD struct {
	_type    string    `json:"TYPE"`
	Market    string    `json:"MARKET"`
	Fromsymbol    string    `json:"FROMSYMBOL"`
	Tosymbol    string    `json:"TOSYMBOL"`
	Flags    string    `json:"FLAGS"`
	Price    float64    `json:"PRICE"`
	Lastupdate    int64    `json:"LASTUPDATE"`
	Lastvolume    float64    `json:"LASTVOLUME"`
	Lastvolumeto    float64    `json:"LASTVOLUMETO"`
	Lasttradeid    string    `json:"LASTTRADEID"`
	Volumeday    float64    `json:"VOLUMEDAY"`
	Volumedayto    float64    `json:"VOLUMEDAYTO"`
	Volume24hour    float64    `json:"VOLUME24HOUR"`
	Volume24hourto    float64    `json:"VOLUME24HOURTO"`
	Openday    float64    `json:"OPENDAY"`
	Highday    float64    `json:"HIGHDAY"`
	Lowday    float64    `json:"LOWDAY"`
	Open24hour    float64    `json:"OPEN24HOUR"`
	High24hour    float64    `json:"HIGH24HOUR"`
	Low24hour    float64    `json:"LOW24HOUR"`
	Lastmarket    string    `json:"LASTMARKET"`
	Change24hour    float64    `json:"CHANGE24HOUR"`
	Changepct24hour    float64    `json:"CHANGEPCT24HOUR"`
	Changeday    float64    `json:"CHANGEDAY"`
	Changepctday    float64    `json:"CHANGEPCTDAY"`
	Supply    float64    `json:"SUPPLY"`
	Mktcap    float64    `json:"MKTCAP"`
	Totalvolume24h    float64    `json:"TOTALVOLUME24H"`
	Totalvolume24hto    float64    `json:"TOTALVOLUME24HTO"`
}

type ETH struct {
	USD	USD	`json:"USD"`
}
type BTC struct {
	USD	USD	`json:"USD"`
}
type LTC struct {
	USD	USD	`json:"USD"`
}
type BCH struct {
	USD	USD	`json:"USD"`
}
