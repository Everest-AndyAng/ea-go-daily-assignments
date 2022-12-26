package song

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
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
