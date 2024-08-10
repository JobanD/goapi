package tools

import (
	"time"
)

type mockDB struct{}

var mockLoginDetais = map[string]LoginDetails{
	"joban": {
		AuthToken: "123ABC",
		Username: "joban",
	},
	"jasleen": {
		AuthToken: "123ABC",
		Username: "jasleen",
	},
	"mandeep": {
		AuthToken: "123ABC",
		Username: "mandeep",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"joban": {
		Coins: 700,
		Username: "joban",
	},
	"jasleen": {
		Coins: 85000,
		Username: "jasleen",
	},
	"mandeep": {
		Coins: 215000,
		Username: "mandeep",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetais[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// Simulate db call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}