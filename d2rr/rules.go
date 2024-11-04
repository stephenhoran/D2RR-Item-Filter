package d2rr

import "strings"

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
	for _, rule := range d.rules.Bases {
		for _, i := range d.items {
			for _, b := range rule.Bases {
				if strings.Contains(i.ENUS, b) {
					i.Replace(b, ColorCode(rule.Color).String()+b)
				}
			}
		}
	}
}
