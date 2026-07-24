package main

import (
	"encoding/json"
	"os"
)

type Data struct {
	Shipments []any `json:"shipments"`
}

func main() {
	data, err := os.ReadFile("shipping_data.json")
	if err != nil {
		panic(err)
	}

	var d Data

	if err := json.Unmarshal(data, &d); err != nil {
		panic(err)
	}

	out, err := json.MarshalIndent(d.Shipments, "", "  ")
	if err != nil {
		panic(err)
	}

	if err := os.WriteFile("shipments.json", out, 0644); err != nil {
		panic(err)
	}
}