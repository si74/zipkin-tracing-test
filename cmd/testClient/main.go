package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {

	var url string
	flag.StringVar(&url, "server_url", "http://localhost:3000", "flag to describe testServer")

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	if strings.Contains(strings.ToLower(string(body)), "ca va") {
		log.Print("Comme ci comme ca.")
	}

}
