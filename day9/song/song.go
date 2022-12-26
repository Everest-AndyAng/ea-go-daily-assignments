package song

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Song struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	Singer string `json:"singer"`
	Writer string `json:"writer"`
	// Release_Date time.Time `json:"release_date"`
}

type SongsController struct {
	Songs []Song `json:"songs"`
}

// Mocked DB
func QuerySongs() []Song {
	var songs []Song
	songs = append(songs, Song{1, "In The End", "Linkin Park", "Linkin Park"})
	songs = append(songs, Song{2, "From The Inside", "Linkin Park", "Linkin Park"})
	return songs
}

func (sa *SongsController) GetSongs(c *gin.Context) {
	c.JSON(http.StatusOK, sa.Songs)

}
