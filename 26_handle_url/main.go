package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://courses.chaicode.com/learn/fast-checkout/250939?priceId=233759"

func main() {
	fmt.Println("URL", myurl)

	result, err := url.Parse(myurl)

	if err != nil {
		panic(err)
	}

	// fmt.Println(result.Scheme)
	// fmt.Println(result.Host)
	// fmt.Println(result.Port())
	// fmt.Println(result.Path)
	// fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("type of url query parameters : %T \n", qparams)

	for _, val := range qparams {
		fmt.Println("Parameter is : ", val)
	}

	// creating own url
	partsofurl := &url.URL{
		Scheme: "https",
		Host:   "courses.chaicode.com",
		Path:   "/cohorts",
	}

	maded_url := partsofurl.String()
	fmt.Println(maded_url)
}
