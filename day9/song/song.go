package song

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Song struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Singer       string    `json:"singer"`
	Writer       string    `json:"writer"`
	Release_Date time.Time `json:"release_date"`
}

type SongsController struct {
	Songs []Song `json:"songs"`
}

// Mocked DB
func QuerySongs() []Song {
	var songs []Song
	songs = append(songs, Song{1, "In The End", "Linkin Park", "Linkin Park", time.Time{}})
	songs = append(songs, Song{2, "From The Inside", "Linkin Park", "Linkin Park", time.Time{}})
	return songs
}

func (sc *SongsController) GetSongs(c *gin.Context) {
	c.JSON(http.StatusOK, sc.Songs)
}

func (sc *SongsController) AddSong(c *gin.Context) {
	song := Song{}
	c.BindJSON(&song)
	sc.Songs = append(sc.Songs, song)
	fmt.Println(sc.Songs)

	c.JSON(http.StatusCreated, song)
}
