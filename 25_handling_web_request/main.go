package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://google.com"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("Type of response : %T", response)
	databytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databytes))
	defer response.Body.Close()

}
