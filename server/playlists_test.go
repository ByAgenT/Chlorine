package server_test

import (
	"chlorine/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMyPlaylistsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/me/playlists", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := server.MyPlaylistsHandler{}
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %d, expect %d", rr.Code, http.StatusOK)
	}
}
