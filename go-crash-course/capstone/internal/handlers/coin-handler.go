package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/SaiKrishna1908/goapi/api"
	"github.com/SaiKrishna1908/goapi/internal/tools"
	"github.com/gorilla/schema"
	"github.com/sirupsen/logrus"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params = api.CoinBalanceParams{}
	var decoder *schema.Decoder = schema.NewDecoder()
	var err error

	err = decoder.Decode(&params, r.URL.Query())

	if err != nil {
		logrus.Error(err)
		api.InteralErrorHandler(w)
		return
	}

	var database *tools.DatabaseInterface
	database, err = tools.NewDatabase()
	if err != nil {
		api.InteralErrorHandler(w)
		return
	}

	var tokenDetails *tools.CoinDetails
	tokenDetails = (*database).GetUserCoins(params.Username)

	if tokenDetails == nil {
		logrus.Error(err)
		api.InteralErrorHandler(w)
		return
	}

	var response = api.CoinBalanceResponse{
		Balance: tokenDetails.Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		logrus.Error(err)
		api.InteralErrorHandler(w)
		return
	}
}
