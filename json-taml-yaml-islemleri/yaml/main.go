package main

import (
	"fmt"
	"log"

	"github.com/go-yaml/yaml"
)

func main() {
	p := Person{"Berk", "Safran", 31}

	// Marshal -> object to X (yaml, json, xml etc.)
	// Unmarshal -> X to object
	y, err := yaml.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(y))
}

// Person ...
type Person struct {
	FirstName string `yaml: "first_name"`
	LastName  string `yaml: "last_name"`
	Age       int    `yaml: "age"`
}
