package main

import (
	"fmt"

	polo "github.com/BWMuller/poloniex"
)

func main() {

	ws, err := polo.NewWSClient()
	if err != nil {
		return
	}

	err = ws.SubscribeTicker()
	if err != nil {
		return
	}

	for {
		fmt.Println(<-ws.Subs["ticker"])
	}
}
