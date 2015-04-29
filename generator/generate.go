package generator

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Key string `yaml:"resource_api_key"`
}

func Generate(filepath string) {
	config, err := openFile(filepath)

	data := getData()
	marshalled, err := yaml.Marshal(&data)
	if err != nil {
		panic("Unable to marshal config to yaml.")
	}

	_, err = config.Write(marshalled)
	if err != nil {
		panic("Unable to write to config file.")
	}
}

func openFile(filepath string) (file *os.File, err error) {
	config, err := os.OpenFile(filepath, os.O_RDWR|os.O_TRUNC, 0660)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func getData() (config Config) {
	var key string

	fmt.Printf("Resource API key > ")
	fmt.Scan(&key)

	config.Key = key

	return config
}
