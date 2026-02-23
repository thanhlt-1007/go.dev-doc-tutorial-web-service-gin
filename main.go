package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{ //nolint:gochecknoglobals
	{
		ID:     "1",
		Title:  "Blue Train",
		Artist: "John Coltrane",
		Price:  56.99, //nolint:mnd
	},
	{
		ID:     "2",
		Title:  "Jeru",
		Artist: "Gerry Mulligan",
		Price:  17.99, //nolint:mnd
	},
	{
		ID:     "3",
		Title:  "Sarah Vaughan and Clifford Brown",
		Artist: "Sarah Vaughan",
		Price:  39.99, //nolint:mnd
	},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum album

	err := ctx.BindJSON(&newAlbum)
	if err != nil {
		return
	}

	albums = append(albums, newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.ID == id {
			ctx.IndentedJSON(http.StatusOK, a)

			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)

	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
