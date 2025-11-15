package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jc0b/fleetdm-listener-go/pkg/handlers"
	"github.com/spf13/viper"
)

type ServerConfig struct {
}

func NewServer() (*http.Server, error) {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.WebhookHandler).Methods("POST")

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler:      router,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return srv, nil
}
