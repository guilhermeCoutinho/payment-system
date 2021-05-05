package http

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/guilhermeCoutinho/payment-system/controller/account"
	"github.com/guilhermeCoutinho/payment-system/controller/transaction"
	"github.com/guilhermeCoutinho/payment-system/dal"
	"github.com/guilhermeCoutinho/payment-system/server/http/wrapper"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type App struct {
	address string
	config  *viper.Viper
	logger  logrus.FieldLogger
	router  *mux.Router
	wrapper *wrapper.HTTPWrapper
}

func NewApp(
	config *viper.Viper,
	logger logrus.FieldLogger,
	dal *dal.DAL,
) (*App, error) {
	app := &App{
		config:  config,
		logger:  logger,
		wrapper: wrapper.NewHTTPWrapper(logger),
	}

	app.buildRoutes(dal)
	app.configureAddress()

	return app, nil
}

func (a *App) configureAddress() {
	a.logger.Info("configuring http address")
	a.address = a.config.GetString("http.address")
}

func (a *App) buildRoutes(dal *dal.DAL) {

	router := mux.NewRouter()
	accountController := account.NewAccount(dal, a.config)
	transactionController := transaction.NewTransaction(dal, a.config)

	a.wrapper.Register(router, "/accounts", accountController)
	a.wrapper.Register(router, "/accounts/{ID}", accountController)

	a.wrapper.Register(router, "/acccounts{ID}/transactions", transactionController)
	a.wrapper.Register(router, "/transactions", transactionController)

	a.router = router
}

func (a *App) ListenAndServe() {
	a.logger.Infof("Starting listening on %s", a.address)
	err := http.ListenAndServe(a.address, a.router)
	if err != nil {
		a.logger.WithError(err).Error("Error on starting server")
		os.Exit(1)
	}
}
