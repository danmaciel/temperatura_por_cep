package main

import (
	"log"
	"net/http"

	"github.com/danmaciel/temperatura_por_cep/internal/infra/handler"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello"))
	})
	http.HandleFunc("/temperature/{cep}", handler.TemperatureHandle)
	log.Fatal(http.ListenAndServe(":8080", nil))

}
