package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldReturnAllSongs(t *testing.T) {
	router := setupGin()
	var songs []Song
	songs = append(songs, Song{1, "In The End", "Linkin Park", "Linkin Park"})
	songs = append(songs, Song{2, "From The Inside", "Linkin Park", "Linkin Park"})
	songsApi := SongsApi{router, songs}
	songsApi.setupGetSongsRouter()
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/api/v1/songs", nil)
	router.ServeHTTP(w, req)

	body := w.Result().Body
	data, _ := io.ReadAll(body)
	var actual []Song
	json.Unmarshal(data, &actual)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, songs, actual)
}
