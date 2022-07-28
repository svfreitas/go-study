package main

import (
	"encoding/json"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Sequence []string          `yaml:"sequence,flow", json:"sequence"` // ,flow convert to flow yaml style
	Mapping  map[string]string `yaml:"mapping,flow", json:"mapping"`
}

func main() {

	data, err := ioutil.ReadFile("./data/file.yaml")

	if err != nil {
		panic(err)
	}

	config := Config{}

	err = yaml.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}

	config.Mapping["earth"] = "brown"
	config.Sequence = append(config.Sequence, "three")

	data, err = json.Marshal(config)
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile("./data/file-converted.json", data, 755)

	if err != nil {
		panic(err)
	}
}
