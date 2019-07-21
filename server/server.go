package server

import(
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// New http server contructor function
func New(mux *mux.Router, serverAddress string) *http.Server {
	svr := &http.Server{
		Addr: serverAddress,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 120 * time.Second,
		Handler: mux,
	}
	return svr
}