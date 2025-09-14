package middleware

import (
	"errors"
	"net/http"

	"github.com/SaiKrishna1908/goapi/api"
	"github.com/SaiKrishna1908/goapi/internal/tools"
	"github.com/sirupsen/logrus"
)

var UnAuthorizedError = errors.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var username string = r.URL.Query().Get("username")
		var token string = r.Header.Get("Authorization")
		var err error

		if username == "" || token == "" {
			logrus.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		var database *tools.DatabaseInterface
		database, err = tools.NewDatabase()
		if err != nil {
			api.InteralErrorHandler(w)
			return
		}

		var loginDetails *tools.LoginDetails

		loginDetails = (*database).GetUserLoginDetails(username)
		logrus.Info(loginDetails)
		logrus.Info(token)
		logrus.Info((*loginDetails).AuthToken)
		if loginDetails == nil || token != (*loginDetails).AuthToken {
			logrus.Error(UnAuthorizedError)
			api.RequestErrorHandler(w, UnAuthorizedError)
			return
		}

		next.ServeHTTP(w, r)
	})
}
