package main

import (
	"github.com/darkkaiser/study-go/webserve06_WebDecoratorHandler/decoHandler"
	"github.com/darkkaiser/study-go/webserve06_WebDecoratorHandler/myapp"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := NewHandler()

	http.ListenAndServe(":3000", mux)
}

func NewHandler() http.Handler {
	h := myapp.NewHandler()
	h = decoHandler.NewDecoHandler(h, logger)
	h = decoHandler.NewDecoHandler(h, logger2)
	return h
}

func logger(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER1] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER1] Completed time:", time.Since(start).Milliseconds())
}

func logger2(w http.ResponseWriter, r *http.Request, h http.Handler) {
	start := time.Now()
	log.Println("[LOGGER2] Started")
	h.ServeHTTP(w, r)
	log.Println("[LOGGER2] Completed time:", time.Since(start).Milliseconds())
}
