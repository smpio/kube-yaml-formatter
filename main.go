package main

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("Error reading stdin: %v\n", err)
	}

	obj := yaml.MapSlice{}

	err = yaml.Unmarshal(data, &obj)
	if err != nil {
		log.Fatalf("Error parsing YAML: %v\n", err)
	}

	data, err = yaml.Marshal(obj)
	if err != nil {
		log.Fatalf("Error formatting YAML: %v\n", err)
	}

	os.Stdout.Write(data)
}
