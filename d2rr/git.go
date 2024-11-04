package d2rr

const tagsUrl string = "https://api.github.com/repos/D2R-Reimagined/d2r-reimagined-mod/tags"
const rawGitPrefix string = "https://raw.githubusercontent.com/D2R-Reimagined/d2r-reimagined-mod/refs/tags/"
const itemFilterSuffix string = "/data/local/lng/strings/item-names.json"

type tagsResponse []tag

type tag struct {
	Name string `json:"name"`
}

func (t tag) String() string {
	return t.Name
}
