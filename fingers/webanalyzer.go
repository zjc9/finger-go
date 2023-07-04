package fingers

import (
	"encoding/json"
	"github.com/kataras/golog"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var log = golog.Child("web")

type WebAnalyzer struct {
	Rules
}

func (w *WebAnalyzer) loadRules(path string) {
	var files = []string{}

	err := filepath.Walk(path, func(p string, info fs.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		// 如果不是文件，则直接返回
		if err != nil {
			log.Error(err)
			return err
		}
		if !strings.HasSuffix(p, "json") {
			return nil
		}
		files = append(files, p)

		return nil
	})

	for _, file := range files {
		log.Info(file)
		bytes, err := os.ReadFile(file)
		if err != nil {
			log.Error(err)
		}
		var rule = &Rule{}

		err = json.Unmarshal(bytes, rule)
		if err != nil {
			log.Error(err)
		}
		//log.Info(rule)
		w.AllRules = append(w.AllRules, *rule)
	}

	//bytes, err := json.Marshal(rules)
	//if err != nil {
	//	log.Error(err)
	//}
	////err = os.WriteFile("all.json", bytes, 0666)
	//err = ioutil.WriteFile("fofa.json", bytes, 0666)
	//if err != nil {
	//	log.Error(err)
	//}

	//log.Info(string(bytes))
	if err != nil {
		log.Error(err)

	}

}

func (w *WebAnalyzer) match(resp *http.Response) {
	for _, rule := range w.AllRules {
		for _, match := range rule.Matches {
			search := match.Search
			switch search {
			case "all":
				SearchAll(resp)
			case "body":
				SearchBody(resp)

			}

		}
	}
}

func SearchHeader(header *http.Header) {

}
func SearchBody(resp *http.Response) {

}

func SearchAll(resp *http.Response) {
	SearchHeader(&resp.Header)
}
