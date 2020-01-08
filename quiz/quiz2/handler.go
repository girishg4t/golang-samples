package urlshort

import (
	"encoding/json"
	"net/http"

	yamlV2 "gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if redURL, ok := pathsToUrls[r.URL.RequestURI()]; ok {
			http.Redirect(w, r, redURL, 301)
		}
		fallback.ServeHTTP(w, r)
	})
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yaml []byte) (dst []map[string]string, err error) {
	var mapYaml []map[string]string
	e := yamlV2.Unmarshal(yaml, &mapYaml)
	return mapYaml, e
}

func buildMap(yaml []map[string]string) map[string]string {
	yamlmap := make(map[string]string, len(yaml))
	for _, m := range yaml {
		yamlmap[m["path"]] = m["url"]
	}
	return yamlmap
}

func parseJSON(jsonBytes []byte) (dst []map[string]string, err error) {
	var mapYaml []map[string]string
	e := json.Unmarshal(jsonBytes, &mapYaml)
	return mapYaml, e
}

//JSONHandler to handle the json
func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseJSON(json)
	if err != nil {
		return nil, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}
