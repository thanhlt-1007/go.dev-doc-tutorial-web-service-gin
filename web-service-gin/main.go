package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

// album represents data about a record album.
type Album struct {
    ID string `json:"id"`
    Title string `json:"title"`
    Artist string `json:"artist"`
    Price float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album {
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbum adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var album Album

    err := c.BindJSON(&album)
    if err != nil {
        return
    }

    albums = append(albums, album)
    c.IndentedJSON(http.StatusCreated, album)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumById(c *gin.Context) {
    id := c.Param("id")

    for _, album := range albums {
        if album.ID == id {
            c.IndentedJSON(http.StatusOK, album)
            return
        }
    }

    c.IndentedJSON(http.StatusNotFound, gin.H { "message": "album not found" })
}

func main() {
    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.GET("/albums/:id", getAlbumById)
    router.POST("/albums", postAlbums)

    router.Run("localhost:8080")
}
