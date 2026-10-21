package main

import (
	"fmt"
	"log"
	"net/http"
	"temp/routers"
)

func main() {
	mux := routers.SetUpRoutes()

	fmt.Println("Starting server at port : 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}
