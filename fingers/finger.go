package fingers

type FingerScanExecInterface interface {
	loadRules(path string)
}

type Rules struct {
	AllRules []Rule `json:"rules"`
}

type Rule struct {
	Name        string    `json:"name"`
	Author      string    `json:"author"`
	Version     string    `json:"version"`
	Description string    `json:"description"`
	Website     string    `json:"website"`
	Matches     []matches `json:"matches"`
	Condition   string    `json:"condition"`
	Implies     string    `json:"implies"`
	Excludes    string    `json:"excludes"`
}
type matches struct {
	Name      string `json:"name"`
	Search    string `json:"search"`
	Regexp    string `json:"regexp"`
	Text      string `json:"text"`
	Version   string `json:"version"`
	Offset    int    `json:"offset"`
	Certainty int    `json:"certainty"`
	Md5       string `json:"md5"`
	Url       string `json:"url"`
	Status    int    `json:"status"`
}
