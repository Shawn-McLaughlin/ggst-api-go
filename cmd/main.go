package main

import (
	"fmt"

	ggst_api "github.com/Shawn-McLaughlin/ggst-api-go"
)

func main() {
	fmt.Println("Beginning GGST API testing...")

	var steamId uint64 = 0
	client := ggst_api.NewClient(steamId)
	loginError := client.Login()
	if loginError != nil {
		panic(loginError)
	}

	client.PrintDetails()

	fmt.Println("Ending GGST API testing")
}
