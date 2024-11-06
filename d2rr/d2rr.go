package d2rr

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type d2rr struct {
	version    string
	projectUrl string
	items      itemResponse
	rules      Rules
}

func (d *d2rr) parseLatestVersion() {
	resp, err := http.Get(tagsUrl)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	versions := tagsResponse{}

	err = json.Unmarshal(body, &versions)
	if err != nil {
		log.Fatalln(err)
	}

	d.version = versions[0].String()
}

func New() d2rr {
	d := d2rr{
		projectUrl: "https://github.com/D2R-Reimagined/d2r-reimagined-mod",
	}

	d.parseLatestVersion()
	d.GetItems()
	d.LoadConfig("rules.yaml")

	return d
}

func (d *d2rr) GetItems() {
	resp, err := http.Get(rawGitPrefix + d.version + itemFilterSuffix)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	items := itemResponse{}

	err = json.Unmarshal(body, &items)
	if err != nil {
		log.Fatalln(err)
	}

	d.items = items
}

func (d *d2rr) Marshal() []byte {
	json, _ := json.MarshalIndent(d.items, "", "\t")
	return json
}

func (d *d2rr) ApplyRules() {
	d.RulesElite()
	d.RulesBases()
	d.RulesJewels()
	d.RulesRemoveNormal()
}
