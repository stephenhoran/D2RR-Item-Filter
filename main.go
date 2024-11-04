package main

import (
	"d2rr-filter/d2rr"
	"os"
)

func main() {
	d2 := d2rr.New()
	d2.ApplyRules()

	err := os.WriteFile("item-names.json", d2.Marshal(), 0644)
	if err != nil {
		panic(err)
	}
}
