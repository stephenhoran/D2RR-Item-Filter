package d2rr

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Rules struct {
	Bases        []Itemsrule       `yaml:"bases"`
	Types        map[string]string `yaml:"types"`
	Jewels       Jewels            `yaml:"jewels"`
	RemoveNormal bool              `yaml:"removeNormal"`
}

type Itemsrule struct {
	Color  string   `yaml:"color"`
	Prefix string   `yaml:"prefix"`
	Bases  []string `yaml:"bases"`
}

type Jewels struct {
	Color string `yaml:"color"`
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
