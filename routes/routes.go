package routes

import(
	"net/http"
	"time"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/gorilla/mux"
)

// Handlers struct
type Handlers struct {
	logger *logrus.Logger
}

// SetupRoutes function
// Register all routers here
func (h *Handlers) SetupRoutes(mux *mux.Router) {
	mux.HandleFunc("/health", h.Logger(h.Health))
	mux.HandleFunc("/", h.Logger(h.Home))
}

// Home Handler function
func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("hit the Home handler with Path : %v", r.URL.Path)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message": "default route",
		"path": fmt.Sprintf("%v", r.URL.Path),
	})
}

// Health Handler function
func (h *Handlers) Health(w http.ResponseWriter, r *http.Request) {
	h.logger.Infof("hit the Health handler with Path : %v", r.URL.Path)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"status": "ok",
		"path": fmt.Sprintf("%v", r.URL.Path),
	})

}

// Logger func
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		defer h.logger.Infof("request processed in %s", time.Now().Sub(startTime))
		next(w, r)
	}
}

// NewHandlers contructor function
func NewHandlers(logger *logrus.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}