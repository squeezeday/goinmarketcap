package coinmarketcapapi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// API URL
const (
	APIURL = "https://api.coinmarketcap.com"
)

// CoinmarketcapAPI API Client
type CoinmarketcapAPI struct {
	client *http.Client
}

// Coin Cryptocurrency
type Coin struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Symbol           string `json:"symbol"`
	Rank             string `json:"rank"`
	PriceUsd         string `json:"price_usd"`
	PriceBtc         string `json:"price_btc"`
	Volume24h        string `json:"24h_volume_usd"`
	MarketCapUsd     string `json:"market_cap_usd"`
	AvailableSupply  string `json:"available_supply"`
	TotalSupply      string `json:"total_supply"`
	MaxSupply        string `json:"max_supply"`
	PercentChange1h  string `json:"percent_change_1h"`
	PercentChange24h string `json:"percent_change_24h"`
	PercentChange7d  string `json:"percent_change_7d"`
	LastUpdated      string `json:"last_updated"`
}

func New() *CoinmarketcapAPI {
	return &CoinmarketcapAPI{http.DefaultClient}
}

// Ticker return Coin array
func (api *CoinmarketcapAPI) Ticker(limit int32) ([]Coin, error) {
	url := fmt.Sprintf("%s/v1/ticker/?limit=%d", APIURL, limit)
	req, err := http.NewRequest("GET", url, strings.NewReader(""))
	res, err := api.client.Do(req)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}
	var jsonData []Coin

	json.Unmarshal(content, &jsonData)
	return jsonData, nil
}
