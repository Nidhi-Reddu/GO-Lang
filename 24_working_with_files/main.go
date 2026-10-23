package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "This is an example code file!"
	file, err := os.Create("./content.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("Length of file : ", length)
	defer file.Close()
	filereading("./content.txt")
}

func filereading(filename string) {
	data_byte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data : ", string(data_byte))
}
