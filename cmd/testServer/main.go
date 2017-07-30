package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "salut! ca va?")
	})

	log.Fatal(http.ListenAndServe(":3000", nil))

}
