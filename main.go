package main

import (
    "os"
    "log"
    "io/ioutil"
    "github.com/ghodss/yaml"
)

func main() {
    data, err := ioutil.ReadAll(os.Stdin)
    if err != nil {
        log.Fatalf("Error reading stdin: %v\n", err)
    }

    var obj interface{}

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
