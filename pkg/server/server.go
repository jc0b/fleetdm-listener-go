package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jc0b/fleetdm-listener-go/pkg/handlers"
	"github.com/spf13/viper"
)

type ServerConfig struct {
}

func NewServer() (*http.Server, error) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.WebhookHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", viper.GetInt("port")),
		Handler:      mux,
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	return srv, nil
}
