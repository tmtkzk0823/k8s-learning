package main

import (
	"fmt",
	"log",
	"net/http"
)

func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r * http.Request){
		fmt.Fprint(w, "Hello, world!")
	})
  log.PrintIn("Staring server on port 8080")
	err := http.ListenAddServe(':8080, nil')
	if err != nil {
		log.Fatal(err)
	}
}

