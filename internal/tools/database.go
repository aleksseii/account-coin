package tools

import logger "github.com/sirupsen/logrus"

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails

	GetUserCoins(username string) *CoinDetails

	SetupDatabase() error
}

type LoginDetails struct {
	username  string
	authToken string
}

type CoinDetails struct {
	coins    int64
	username string
}

func NewDatabse() (*DatabaseInterface, error) {
	var database DatabaseInterface = &MockDB{}
	err := database.SetupDatabase()
	if err != nil {
		logger.Fatalf("Failed to setup database, err: %v", err)
	}
	return &database, err
}
