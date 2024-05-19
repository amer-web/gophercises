package urlshort

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"net/http"
)

type YamlRedirect struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if urlShort, ok := pathsToUrls[r.URL.Path]; ok {
			fmt.Fprintf(w, "redirct to %v", urlShort)
			//http.Redirect(w, r, urlShort, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}

}

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {

	var yamlR []YamlRedirect

	if er := yaml.Unmarshal(yml, &yamlR); er != nil {
		fmt.Println(yamlR, er.Error())
	}
	var myMap = make(map[string]string)
	for _, y := range yamlR {
		myMap[y.Path] = y.URL
	}
	fmt.Println(myMap)

	return MapHandler(myMap, fallback), nil

}
