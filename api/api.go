package api

import (
	"encoding/json"
	"net/http"
)


//coin balance params
type CoinBalanceParams struct{
	Username string
}

//coin balance Response
type CoinBalanceResponse struct{
	Code int

	Balance int64
}


// Error Response
type Error struct{
	Code int
	Message string
}