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

	config := generator.Config{TargetName: "", EmojiNum: 4, PunctiuationLevel: 0}

	h := http.NewServeMux()
	h.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		text, err := generator.Start(config)
		if err == nil {
			fmt.Fprintf(w, "<html><body><p style=\"font-size:200%%\">%s</p><footer>Inspired by <a href=\"https://reverent-shirley-368990.netlify.com/\">ojichat-web</a>, fork me on <a href=\"https://github.com/zunda/ojichat-plaintext\">GitHub</a></footer></body></html>", text)
		} else {
			log.Fatal(err)
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
	})

	log.Println("Listening at port " + port)
	err := http.ListenAndServe(":"+port, h)
	log.Fatal(err)
}
