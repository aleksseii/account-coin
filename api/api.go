package api

import (
	"encoding/json"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

type CoinBalanceParams struct {
	Username string
}

type CoinBalanceResponse struct {
	Code    uint16
	Balance int64
}

type Error struct {
	code    uint16
	message string
}

func writeError(responseWriter http.ResponseWriter, code uint16, message string) {
	response := Error{
		code:    code,
		message: message,
	}
	responseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
	responseWriter.WriteHeader(int(code))

	err := json.NewEncoder(responseWriter).Encode(response)
	if err != nil {
		logger.Errorf("Failed to write response, err: %v", err)
	}
}

var (
	RequestErrorHandler = func(responseWriter http.ResponseWriter, err error) {
		writeError(responseWriter, http.StatusBadRequest, err.Error())
	}
	InternalErrorHandler = func(responseWriter http.ResponseWriter) {
		writeError(responseWriter, http.StatusInternalServerError, "Unexpected internal error occurred")
	}
)
