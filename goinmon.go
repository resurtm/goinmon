package main

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"
	"github.com/olekukonko/tablewriter"
	"os"
	_ "github.com/dimiro1/banner/autoload"
)

const apiURL = "https://api.coinmarketcap.com/v1/ticker/?limit=10&convert=USD"

type Currency struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
	Rank   int    `json:"rank,string"`

	PriceUSD float64 `json:"price_usd,string"`
	PriceBTC float64 `json:"price_btc,string"`

	Last24HVolumeUSD float64 `json:"24h_volume_usd,string"`
	MarketCapUSD     float64 `json:"market_cap_usd,string"`

	AvailableSupply float64 `json:"available_supply,string"`
	TotalSupply     float64 `json:"total_supply,string"`
	MaxSupply       float64 `json:"max_supply,string"`

	PercentChange1H  float64 `json:"percent_change_1h,string"`
	PercentChange24H float64 `json:"percent_change_24h,string"`
	PercentChange7D  float64 `json:"percent_change_7d,string"`

	LastUpdated int `json:"last_updated,string"`
}

type Currencies struct {
	Currencies []Currency
}

func (c *Currencies) LoadData() error {
	resp, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &c.Currencies)
}

func main() {
	currencies := Currencies{}
	currencies.LoadData()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Rank",
		"Coin",
		"Price (USD)",
		"Change (24H)",
		"Change (1H)",
		"Market Cap (USD)",
	})
	for _, currency := range currencies.Currencies {
		table.Append([]string{
			strconv.Itoa(currency.Rank),
			currency.Symbol,
			strconv.FormatFloat(currency.PriceUSD, 'f', 3, 64),
			strconv.FormatFloat(currency.PercentChange24H, 'f', 2, 64),
			strconv.FormatFloat(currency.PercentChange1H, 'f', 2, 64),
			strconv.FormatFloat(currency.MarketCapUSD, 'f', 2, 64),
		})
	}
	table.Render()
}
