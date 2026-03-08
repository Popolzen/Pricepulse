package main

import (
	"net/http"

	"github.com/popolzen/pricepulse/internal/handler"
)

func main() {
	http.HandleFunc(`/`, handler.Health)
	http.ListenAndServe("localhost:8080", nil)

}
