package main

import (
	"encoding/json"
	"io/ioutil"
)

type Register struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	ID        int    `json:"id"`
}

func main() {

	data, err := ioutil.ReadFile("./data/file.json")

	if err != nil {
		panic(err)
	}

	registers := []Register{}

	err = json.Unmarshal(data, &registers)
	if err != nil {
		panic(err)
	}

	registers = append(registers, Register{FirstName: "Jane", LastName: "Doe", ID: 7890})

	data, err = json.Marshal(registers)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./data/file-processed.json", data, 755)

	if err != nil {
		panic(err)
	}
}
