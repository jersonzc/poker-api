package poker_api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	t.Run("returns Jerson's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Jerson", nil)
		response := httptest.NewRecorder()

		PlayerServer(request, response)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
