package d2rr

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Rules struct {
	Bases []Itemsrule       `yaml:"bases"`
	Types map[string]string `yaml:"types"`
}

type Itemsrule struct {
	Color string   `yaml:"color"`
	Bases []string `yaml:"bases"`
}

func (d *d2rr) LoadConfig(f string) {
	var rules Rules

	yamlfile, err := os.ReadFile(f)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(yamlfile, &rules)
	if err != nil {
		panic(err)
	}

	d.rules = rules
}
