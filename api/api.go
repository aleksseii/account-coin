package api

type CoinBalanceParams struct {
	username string
}

type CoinBalanceResponse struct {
	code    uint16
	balance int64
}

type Error struct {
	code    uint16
	message string
}
