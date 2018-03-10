package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Package struct {
	Info `json:"info"`
}

type Info struct {
	Author       string
	Author_email string
	Download_url string
	Version      string
}

func main() {
	resp, err := http.Get("https://pypi.python.org/pypi/yaspin/json")
	if err != nil {
		log.Fatal(err)
	}

	bs, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Printf("%s \n", bs)

	var pack Package
	json.Unmarshal(bs, &pack)

	fmt.Printf("%+v\n", pack)
}
