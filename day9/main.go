package main

import (
	"encoding/json"
	"fmt"
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

type SongsApi struct {
	Router *gin.Engine
	Songs  []Song `json:"songs"`
}

func main() {
	router := setupGin()
	songsApi := SongsApi{router, setupSongs()}
	songsApi.setupGetSongsRouter()
	router.Run(":8080")
}

func setupGin() *gin.Engine {
	return gin.Default()
}

func setupSongs() []Song {
	var songs []Song
	songs = append(songs, Song{1, "In The End", "Linkin Park", "Linkin Park"})
	songs = append(songs, Song{2, "From The Inside", "Linkin Park", "Linkin Park"})
	return songs
}

func (sa *SongsApi) setupRouter() {
	sa.setupGetSongsRouter()
}

func (sa *SongsApi) setupGetSongsRouter() {
	result, _ := json.Marshal(sa.Songs)
	fmt.Println(string(result))
	sa.Router.GET("/api/v1/songs", func(c *gin.Context) {
		c.JSON(http.StatusOK, sa.Songs)
	})
}
