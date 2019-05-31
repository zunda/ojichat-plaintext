package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/greymd/ojichat/generator"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

  config := generator.Config{}

	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    text, err = generator.Start(config)
    if err == nil {
      fmt.Fprintf(w, text)
    } else {
      log.Error(err)
      w.WriteHeder(http.StatusServiceUnavailable)
    }
	})

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
