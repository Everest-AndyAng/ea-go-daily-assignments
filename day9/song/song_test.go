package song

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var router *gin.Engine

func init() {
	router = gin.New()
}

func TestShouldReturnAllSongsWhenGettingAllSongs(t *testing.T) {
	url := "/api/v1/songs"
	songs := QuerySongs()
	controller := SongsController{songs}

	router.GET(url, controller.GetSongs)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, url, nil)
	router.ServeHTTP(w, req)

	body := w.Result().Body
	data, _ := io.ReadAll(body)
	var actual []Song
	json.Unmarshal(data, &actual)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, songs, actual)
}

func TestShouldAddSongToSongListWhenCreatingNewSong(t *testing.T) {
	url := "/api/v1/songs"
	songs := QuerySongs()
	newSong := Song{3, "Numb", "Linkin Park", "Linkin Park", time.Time{}}
	songToAdd, _ := json.Marshal(newSong)
	controller := SongsController{songs}

	router.POST(url, controller.AddSong)

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, url, strings.NewReader(string(songToAdd)))
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Containsf(t, controller.Songs, newSong, "New song not added")
}
