package cyoa_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/FedericoBarberon/Go-Exercises/cyoa"
)

func TestHandler(t *testing.T) {
	t.Run("shows intro arc at /", func(t *testing.T) {
		handler, err := cyoa.NewHandler(exampleBook)

		assertNoError(t, err)

		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, "text/html")
		assertStringContains(t, response.Body.String(), "<h1>test</h1>")
	})
	t.Run("shows arc 'arc 1' at /arc-1", func(t *testing.T) {
		handler, err := cyoa.NewHandler(exampleBook)

		assertNoError(t, err)

		request, _ := http.NewRequest(http.MethodGet, "/arc-1", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, "text/html")
		assertStringContains(t, response.Body.String(), "<h1>arc 1</h1>")
	})
	t.Run("shows 404 when hit a non-existing arc", func(t *testing.T) {
		handler, err := cyoa.NewHandler(exampleBook)

		assertNoError(t, err)

		request, _ := http.NewRequest(http.MethodGet, "/non-existing", nil)
		response := httptest.NewRecorder()

		handler.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
		assertContentType(t, response, "text/html")
		assertStringContains(t, response.Body.String(), "<h1>404 - Arc Not Found</h1>")
	})
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("expected status %d but got %d", want, got)
	}
}

func assertContentType(t testing.TB, res http.ResponseWriter, want string) {
	t.Helper()
	got := res.Header().Get("content-type")

	if got != want {
		t.Errorf("expected a content-type of %s but got %s", want, got)
	}
}

func assertStringContains(t testing.TB, str, substr string) {
	t.Helper()
	if !strings.Contains(str, substr) {
		t.Errorf("expected %s to have %s but didnt have it", str, substr)
	}
}
