package d2rr

import (
	"strings"
)

type itemResponse []*item

type item struct {
	ID   int    `json:"id"`
	Key  string `json:"Key"`
	ENUS string `json:"enUS"`
	ZHTW string `json:"zhTW"`
	DEDE string `json:"deDE"`
	ESES string `json:"esES"`
	FRFR string `json:"frFR"`
	ITIT string `json:"itIT"`
	KOKR string `json:"koKR"`
	PLPL string `json:"plPL"`
	ESMX string `json:"esMX"`
	JAJP string `json:"jaJP"`
	PTBR string `json:"ptBR"`
	RURU string `json:"ruRU"`
	ZHCN string `json:"zhCN"`
}

func (i *item) Replace(filter string, replace string) {
	i.ENUS = strings.Replace(i.ENUS, filter, replace, 1)
}

func (d *d2rr) ReturnItems() itemResponse {
	return d.items
}

func (d *d2rr) FindItem(i string) []item {
	var items []item
	for _, v := range d.items {
		if strings.Contains(v.ENUS, i) {
			items = append(items, *v)
		}
	}

	return items
}
