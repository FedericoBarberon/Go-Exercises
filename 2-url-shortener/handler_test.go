package urlshort_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FedericoBarberon/urlshort"
)

func TestMapHandler(t *testing.T) {
	defaultMux := http.NewServeMux()
	defaultMux.HandleFunc("/notMapped", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	})
	defaultMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "default mux")
	})

	paths := map[string]string{
		"/hello":     "/bye",
		"/a":         "/b",
		"/notMapped": "",
	}

	redirectCases := []struct {
		path       string
		redirectTo string
	}{
		{
			path:       "/hello",
			redirectTo: "/bye",
		},
		{
			path:       "/hello",
			redirectTo: "/bye",
		},
	}

	noRedirectCases := []struct {
		path    string
		bodyStr string
	}{
		{
			path:    "/notMapped",
			bodyStr: "test",
		},
		{
			path:    "/",
			bodyStr: "default mux",
		},
	}

	for _, testCase := range redirectCases {
		t.Run(fmt.Sprintf("redirects %s to %s", testCase.path, testCase.redirectTo), func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, testCase.path, nil)
			response := httptest.NewRecorder()

			handler := urlshort.MapHandler(paths, defaultMux)

			handler.ServeHTTP(response, request)

			statusGot := response.Code
			statusWanted := http.StatusSeeOther

			if statusGot != statusWanted {
				t.Errorf("expected status %d but got %d", statusWanted, statusGot)
			}

			urlObj, _ := response.Result().Location()
			urlGot := urlObj.String()

			if urlGot != testCase.redirectTo {
				t.Errorf("expected redirect to %s but got %s", testCase.redirectTo, urlGot)
			}
		})
	}

	for _, testCase := range noRedirectCases {
		t.Run("use default mux on "+testCase.path, func(t *testing.T) {
			request, _ := http.NewRequest(http.MethodGet, testCase.path, nil)
			response := httptest.NewRecorder()

			handler := urlshort.MapHandler(paths, defaultMux)

			handler.ServeHTTP(response, request)

			got := response.Body.String()

			if got != testCase.bodyStr {
				t.Errorf("expected %s but got %s", testCase.bodyStr, got)
			}
		})
	}
}
