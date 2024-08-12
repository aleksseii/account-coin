package handler

import (
	"account-coin/api"
	"account-coin/internal/tools"
	"encoding/json"
	"github.com/gorilla/schema"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

func GetCoinBalance(responseWriter http.ResponseWriter, request *http.Request) {

	params := api.CoinBalanceParams{}
	decoder := schema.NewDecoder()

	err := decoder.Decode(&params, request.URL.Query())
	if err != nil {
		logger.Errorf("Failed to decode coin balance params, err: %v", err)
		api.InternalErrorHandler(responseWriter)
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalErrorHandler(responseWriter)
		return
	}

	err = (*database).SetupDatabase()
	if err != nil {
		api.InternalErrorHandler(responseWriter)
		return
	}

	coinDetails := (*database).GetUserCoins(params.Username)
	if coinDetails == nil {
		logger.Errorf("Failed to get coin details for user: %s", params.Username)
		api.InternalErrorHandler(responseWriter)
		return
	}

	response := api.CoinBalanceResponse{
		Code:    http.StatusOK,
		Balance: (*coinDetails).Coins,
	}

	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	err = json.NewEncoder(responseWriter).Encode(response)
	if err != nil {
		logger.Errorf("Failed to write coin balance response, err: %v", err)
		api.InternalErrorHandler(responseWriter)
		return
	}
}
