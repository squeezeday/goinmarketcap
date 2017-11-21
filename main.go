package main

import (
	"fmt"
	"strconv"

	"./coinmarketcap-api"
)

func main() {

	api := coinmarketcapapi.New()
	res, err := api.Ticker(10)
	if err != nil {
		fmt.Errorf("Error: %s", err.Error())
	}

	fmt.Printf("%-5s %8s\n", "coin", "price")
	for _, coin := range res {
		usd, _ := strconv.ParseFloat(coin.PriceUsd, 64)
		fmt.Printf("%-5s %8.2f\n", coin.Symbol, usd)
	}
}
