package urlshort_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/FedericoBarberon/Go-Exercises/urlshort"
)

func TestMapHandler(t *testing.T) {
	fallback := fallbackMux()

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
			path:       "/a",
			redirectTo: "/b",
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

	handler := urlshort.MapHandler(paths, fallback)

	for _, testCase := range redirectCases {
		t.Run(fmt.Sprintf("redirects %s to %s", testCase.path, testCase.redirectTo), func(t *testing.T) {
			assertRedirect(t, handler, testCase.path, testCase.redirectTo)
		})
	}

	for _, testCase := range noRedirectCases {
		t.Run("use fallback on "+testCase.path, func(t *testing.T) {
			assertBody(t, handler, testCase.path, testCase.bodyStr)
		})
	}
}

func TestYAMLHandler(t *testing.T) {
	t.Run("valid YAML", func(t *testing.T) {
		fallback := fallbackMux()

		pathsYAML := `
- path: /hello
  url: /bye
- path: /a
  url: /b`

		redirectCases := []struct {
			path       string
			redirectTo string
		}{
			{
				path:       "/hello",
				redirectTo: "/bye",
			},
			{
				path:       "/a",
				redirectTo: "/b",
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

		handler, _ := urlshort.YAMLHandler([]byte(pathsYAML), fallback)

		for _, testCase := range redirectCases {
			t.Run(fmt.Sprintf("redirects %s to %s", testCase.path, testCase.redirectTo), func(t *testing.T) {
				assertRedirect(t, handler, testCase.path, testCase.redirectTo)
			})
		}

		for _, testCase := range noRedirectCases {
			t.Run("use fallback on "+testCase.path, func(t *testing.T) {
				assertBody(t, handler, testCase.path, testCase.bodyStr)
			})
		}
	})
	t.Run("invalid YAML", func(t *testing.T) {
		_, err := urlshort.YAMLHandler([]byte("invalid yaml"), nil)
		assertError(t, err, urlshort.ErrInvalidYAML)
	})
}

func fallbackMux() http.Handler {
	fallback := http.NewServeMux()
	fallback.HandleFunc("/notMapped", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "test")
	})
	fallback.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "default mux")
	})

	return fallback
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("Expected an error but didn't get one")
	}

	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}
}

func assertRedirect(t testing.TB, handler http.Handler, path, redirectTo string) {
	t.Helper()

	request, _ := http.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	statusGot := response.Code
	statusWanted := http.StatusSeeOther

	if statusGot != statusWanted {
		t.Errorf("expected status %d but got %d", statusWanted, statusGot)
	}

	urlObj, err := response.Result().Location()

	if err != nil {
		t.Fatal("no location provided ", err)
	}

	urlGot := urlObj.String()

	if urlGot != redirectTo {
		t.Errorf("expected redirect to %s but got %s", redirectTo, urlGot)
	}
}

func assertBody(t testing.TB, handler http.Handler, path, bodyStr string) {
	t.Helper()

	request, _ := http.NewRequest(http.MethodGet, path, nil)
	response := httptest.NewRecorder()

	handler.ServeHTTP(response, request)

	got := response.Body.String()

	if got != bodyStr {
		t.Errorf("expected %s but got %s", bodyStr, got)
	}
}
