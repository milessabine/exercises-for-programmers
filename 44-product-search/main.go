package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/milessabine/prompt"
)

type Product struct {
	Name     string
	Price    float64 `json:"cost"`
	Quantity int
}

type Base struct {
	Products []Product
}

func main() {

	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal("Unable to open file", err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	var b Base
	err = dec.Decode(&b)
	if err != nil {
		log.Fatal("Unable to decode", err)
	}
	p := prompt.New()
	name := p.Scan("What is the product name?")

	for _, p := range b.Products {
		if p.Name == name {
			fmt.Printf("Name: %s\n", p.Name)
			fmt.Printf("Price: %0.2f\n", p.Price)
			fmt.Printf("Quantity: %d\n", p.Quantity)
		}
	}
}
