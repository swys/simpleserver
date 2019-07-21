package main

import (
	"simpleserver/routes"
	"simpleserver/server"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

var (
	// ListenAddr the http server will listen on (environment variable)
	ListenAddr = os.Getenv("LISTEN_ADDR")
)

func main() {
	// init logger
	log := logrus.New()

	// setup logger
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.InfoLevel)
	log.SetReportCaller(true)

	// init handlers
	h := routes.NewHandlers(log)

	// init mux router
	mux := mux.NewRouter()

	// setup routes
	h.SetupRoutes(mux)

	// init http server
	svr := server.New(mux, ListenAddr)

	log.Info("server is starting up...")

	// Listen and Serve
	log.Fatal(svr.ListenAndServe())
}
