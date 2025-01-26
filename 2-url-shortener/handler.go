package urlshort

import (
	"errors"
	"net/http"

	"gopkg.in/yaml.v3"
)

type PathsToUrls map[string]string

type redirectData struct {
	Path string `yaml:"path"`
	Url  string `yaml:"url"`
}

var ErrInvalidYAML = errors.New("invalid yaml")

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls PathsToUrls, fallback http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		redirectTo, ok := pathsToUrls[r.URL.String()]

		if !ok || redirectTo == "" {
			fallback.ServeHTTP(w, r)
			return
		}

		http.Redirect(w, r, redirectTo, http.StatusSeeOther)
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
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.

func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	redirectData, err := parseYAML(yml)

	if err != nil {
		return nil, err
	}

	return MapHandler(buildMap(redirectData), fallback), nil
}

func parseYAML(yml []byte) ([]redirectData, error) {
	redirectData := []redirectData{}

	err := yaml.Unmarshal(yml, &redirectData)

	if err != nil {
		return nil, ErrInvalidYAML
	}

	return redirectData, nil
}

func buildMap(redirectData []redirectData) PathsToUrls {
	pathsToUrls := make(PathsToUrls)

	for _, redirect := range redirectData {
		pathsToUrls[redirect.Path] = redirect.Url
	}

	return pathsToUrls
}
