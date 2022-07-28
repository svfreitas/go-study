package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	data, err := ioutil.ReadFile("./data/file.txt")

	if err != nil {
		panic(err)
	}

	fixedStr := strings.Replace(string(data), " Trek", " Wars", 1)

	fmt.Println(fixedStr)

	err = ioutil.WriteFile("./data/file-processed.txt", []byte(fixedStr), 755)

	if err != nil {
		panic(err)
	}
}
