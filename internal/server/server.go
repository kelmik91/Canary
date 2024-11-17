package server

import (
	"main/internal/handlers"
	"main/internal/logger"
	"net/http"
)

func Run(host string) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/canary-client", handlers.GetClient)
	mux.HandleFunc("/canary", handlers.PostHandlerCanary)

	if err := http.ListenAndServe(host, mux); err != nil {
		logger.Error(err)
		return err
	}
	return nil
}
