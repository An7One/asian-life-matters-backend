package main

import (
	"net/http"

	"github.com/zea7ot/web_api_aeyesafe/controller"
)

func main() {
	controller.RegisterRoutes()

	http.ListenAndServe("localhost:8080", nil)
}
