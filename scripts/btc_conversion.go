package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

// Struct to parse the JSON response from CoinGecko
type PriceResponse struct {
	Bitcoin struct {
		Usd float64 `json:"usd"`
	} `json:"bitcoin"`
}

func getBTCPriceInUSD() (float64, error) {
	// Make the request to CoinGecko API for Bitcoin price
	resp, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd")
	if err != nil {
		return 0, fmt.Errorf("failed to fetch BTC price: %w", err)
	}
	defer resp.Body.Close()

	// Parse the JSON response
	var priceResponse PriceResponse
	if err := json.NewDecoder(resp.Body).Decode(&priceResponse); err != nil {
		return 0, fmt.Errorf("failed to parse BTC price response: %w", err)
	}

	return priceResponse.Bitcoin.Usd, nil
}

func main() {
	// Check for a USD amount argument
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run btc_conversion.go <USD amount>")
		return
	}

	// Parse the USD amount from the command line
	usdAmount, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		fmt.Println("Invalid USD amount. Please enter a valid number.")
		return
	}

	// Get the current BTC price in USD
	btcPrice, err := getBTCPriceInUSD()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Calculate the BTC amount for the specified USD amount
	btcAmount := usdAmount / btcPrice
	fmt.Printf("The current price of 1 BTC is $%.2f USD\n", btcPrice)
	fmt.Printf("For $%.2f USD, you would get approximately %.8f BTC\n", usdAmount, btcAmount)
}
