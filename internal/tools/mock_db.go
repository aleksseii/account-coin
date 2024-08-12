package tools

import "time"

type MockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		username:  "alex",
		authToken: "fsdga1343551423gasf",
	},
	"jason": {
		username:  "jason",
		authToken: "daskhgfi2384192379ahskdf",
	},
	"marie": {
		username:  "marie",
		authToken: "sdkjhfka27y34i1ykabsdfk",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		username: "alex",
		coins:    62123,
	},
	"jason": {
		username: "jason",
		coins:    13123123,
	},
	"marie": {
		username: "marie",
		coins:    123,
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
