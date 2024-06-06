package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/chronark/flightcontrol/pkg/env"

	"github.com/chronark/flightcontrol/pkg/logging"
)

var (
	port = env.String("PORT", "8080")
)

func main() {

	logConfig := &logging.Config{
		Debug:  os.Getenv("DEBUG") != "",
		Writer: []io.Writer{},
	}

	logger := logging.New(logConfig)

	logger.Info().Strs("env", os.Environ()).Msg("Environment variables")

	http.HandleFunc("/v1/liveness", func(w http.ResponseWriter, r *http.Request) {
		logger.Info().Str("path", r.URL.Path).Str("userAgent", r.UserAgent()).Msg("Liveness check")
		_, err := w.Write([]byte(fmt.Sprintf("%d", time.Now().UnixMilli())))
		if err != nil {

			logger.Err(err).Msg("Error writing response")
		}
	})

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	logger.Info().Str("port", port).Msg("Listening")
	err := server.ListenAndServe()
	if err != nil {
		logger.Err(err).Msg("Error listening")
		os.Exit(1)
	}
}
