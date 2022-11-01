package pages

import (
	"encoding/json"
	"errors"
	"math"
	"net/http"
	"time"
)

type Fetch struct {
	data *[]CoinGeko
	time time.Time
}

func NewFetch() Fetch {
	next := time.Now().Add(10 * time.Second)
	json, _ := getJSON() // FIXME handle err
	return Fetch{data: &json, time: next}
}

func (f *Fetch) Refresh() bool {
	now := time.Now()

	if f.time.Before(now) {
		f.time = time.Now().Add(10 * time.Second)
		return true
	}
	return false
}

func (f *Fetch) Update() {
	json, err := getJSON()

	if err == nil {
		*f.data = json
	}
}

func getJSON() ([]CoinGeko, error) {
	url := "https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=100&page=1&sparkline=false"

	res, err := http.Get(url)
	if err != nil {
		return make([]CoinGeko, 0), errors.New("unable to load page")
	}

	var cg []CoinGeko
	err = json.NewDecoder(res.Body).Decode(&cg)
	if err != nil {
		return make([]CoinGeko, 0), err
	}

	parseJSON(&cg)

	return cg, nil
}

func parseJSON(cg *[]CoinGeko) {
	json := *cg
	for i := range json {
		price := json[i].CurrentPrice
		change := json[i].PriceChange24H

		json[i].CurrentPrice = roundFloat(price, 4)
		json[i].PriceChangePercentage24H = roundFloat(change, 2)
	}
}

// roundFloat rounds the val to the precision's decimal place
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

type CoinGeko struct {
	ID                           string   `json:"id"`
	Symbol                       string   `json:"symbol"`
	Name                         string   `json:"name"`
	Image                        string   `json:"image"`
	CurrentPrice                 float64  `json:"current_price"`
	MarketCap                    int64    `json:"market_cap"`
	MarketCapRank                int64    `json:"market_cap_rank"`
	FullyDilutedValuation        *int64   `json:"fully_diluted_valuation"`
	TotalVolume                  float64  `json:"total_volume"`
	High24H                      float64  `json:"high_24h"`
	Low24H                       float64  `json:"low_24h"`
	PriceChange24H               float64  `json:"price_change_24h"`
	PriceChangePercentage24H     float64  `json:"price_change_percentage_24h"`
	MarketCapChange24H           float64  `json:"market_cap_change_24h"`
	MarketCapChangePercentage24H float64  `json:"market_cap_change_percentage_24h"`
	CirculatingSupply            float64  `json:"circulating_supply"`
	TotalSupply                  *float64 `json:"total_supply"`
	MaxSupply                    *float64 `json:"max_supply"`
	Ath                          float64  `json:"ath"`
	AthChangePercentage          float64  `json:"ath_change_percentage"`
	AthDate                      string   `json:"ath_date"`
	Atl                          float64  `json:"atl"`
	AtlChangePercentage          float64  `json:"atl_change_percentage"`
	AtlDate                      string   `json:"atl_date"`
	Roi                          *Roi     `json:"roi"`
	LastUpdated                  string   `json:"last_updated"`
}

type Roi struct {
	Times      float64  `json:"times"`
	Currency   Currency `json:"currency"`
	Percentage float64  `json:"percentage"`
}

type Currency string

const (
	Btc Currency = "btc"
	Eth Currency = "eth"
	Usd Currency = "usd"
)
