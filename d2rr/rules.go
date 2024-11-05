package d2rr

import (
	"strings"
)

func (d *d2rr) RulesElite() {
	val, ok := d.rules.Types["elite"]
	if ok {
		c := ColorCode(val)

		for _, i := range d.items {
			i.Replace("ÿcN[ÿc4EÿcN]ÿc5", "ÿcN["+c.String()+"EÿcN]ÿc5")
		}
	}
}

func (d *d2rr) RulesBases() {
	var prefix string

	for _, rule := range d.rules.Bases {
		if rule.Prefix != "" {
			prefix = rule.Prefix
		}
		for _, i := range d.items {
			for _, b := range rule.Bases {
				if strings.Contains(i.ENUS, b) {
					i.Replace(b, prefix+ColorCode(rule.Color).String()+b)
				}
			}
		}
	}
}

func (d *d2rr) RulesJewels() {
	if d.rules.Jewels.Color != "" {
		c := ColorCode(d.rules.Jewels.Color)

		for _, i := range d.items {
			i.Replace("Jewelÿc2*ÿc0", "* "+c.String()+"Jewelÿc2*ÿc0")
		}
	}
}
