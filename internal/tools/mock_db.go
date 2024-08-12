package tools

import "time"

type MockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		Username:  "alex",
		AuthToken: "fsdga1343551423gasf",
	},
	"jason": {
		Username:  "jason",
		AuthToken: "daskhgfi2384192379ahskdf",
	},
	"marie": {
		Username:  "marie",
		AuthToken: "sdkjhfka27y34i1ykabsdfk",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Username: "alex",
		Coins:    62123,
	},
	"jason": {
		Username: "jason",
		Coins:    13123123,
	},
	"marie": {
		Username: "marie",
		Coins:    123,
	},
}

func (db *MockDB) GetUserLoginDetails(username string) *LoginDetails {
	time.Sleep(1 * time.Second)

	loginDetails, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}
	return &loginDetails
}

func (db *MockDB) GetUserCoins(username string) *CoinDetails {
	time.Sleep(1 * time.Second)

	coinDetails, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}
	return &coinDetails
}

func (db *MockDB) SetupDatabase() error {
	return nil
}
