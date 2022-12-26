package main

import (
	"music/song"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupGin()
	songsController := setupSongsController()

	router.GET("/api/v1/songs", songsController.GetSongs)
	router.Run(":8080")
}

func setupGin() *gin.Engine {
	return gin.New()
}

func setupSongsController() song.SongsController {
	return song.SongsController{song.QuerySongs()}
}
