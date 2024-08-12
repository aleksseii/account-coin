package middleware

import (
	"account-coin/api"
	"account-coin/internal/tools"
	"errors"
	logger "github.com/sirupsen/logrus"
	"net/http"
)

var unauthorizedError = errors.New("invalid username or token")

func Authorization(nextHandler http.Handler) http.Handler {
	return http.HandlerFunc(
		func(responseWriter http.ResponseWriter, request *http.Request) {
			var err error
			username := request.URL.Query().Get("username")
			incomingToken := request.Header.Get("Authorization")

			if username == "" || incomingToken == "" {
				logger.Error(unauthorizedError)
				api.RequestErrorHandler(responseWriter, unauthorizedError)
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

			loginDetails := (*database).GetUserLoginDetails(username)
			if loginDetails == nil || incomingToken != (*loginDetails).AuthToken {
				logger.Errorf("unknown user or user with invalid token, err: %v", err)
				api.RequestErrorHandler(responseWriter, unauthorizedError)
				return
			}

			nextHandler.ServeHTTP(responseWriter, request)
		},
	)
}
