package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	PORT string
)

func init() {

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "3000"
	}
}

func main() {

	http.HandleFunc("/v1/liveness", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("ok"))
	})

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	fmt.Println("listening on", PORT)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil); err != nil {
		log.Fatalln("error starting server:", err)
	}

	<-ctx.Done()
	log.Println("shutting down")

}
