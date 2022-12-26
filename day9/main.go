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
	router := gin.New()

	router.Use(gin.Recovery())

	return router
}

func setupSongsController() song.SongsController {
	return song.SongsController{song.QuerySongs()}
}
