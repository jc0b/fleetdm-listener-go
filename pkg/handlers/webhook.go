package handlers

import (
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	log.Infof("[webhook] - %s - %s - %s - %s", r.Method, r.RemoteAddr, r.RequestURI, r.Header.Get("User-Agent"))
	if r.Method != "POST" {
		log.Info("Rejected invalid request method")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	if r.Header.Get("Content-Type") != "application/json" && r.Header.Get("content-type") != "application/json" {
		log.Info("Rejected invalid content-type")
		http.Error(w, "Invalid request content-type", http.StatusBadRequest)
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	if len(requestBody) <= 0 {
		http.Error(w, "Invalid Request Body", http.StatusBadRequest)
		return
	}
	//TODO: You can do your thing here
	log.Infof(string(requestBody))
}
