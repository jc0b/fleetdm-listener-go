package handlers

import (
	"fmt"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("[webhook] - %s - %s - %s", r.RemoteAddr, r.RequestURI, r.Header.Get("User-Agent"))
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if r.Header.Get("Content-Type") != "application/json" && r.Header.Get("content-type") != "application/json" {
		http.Error(w, "Invalid request content-type", http.StatusBadRequest)
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	fmt.Println(string(requestBody))
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	if len(requestBody) <= 0 {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
}
