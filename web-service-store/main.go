package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)

	err := router.Run("localhost:8080")
	if err != nil {
		return
	}
}

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
	{ID: "4", Title: "Abbey Road", Artist: "The Beatles", Price: 99.99},
	{ID: "5", Title: "OK Computer", Artist: "Radiohead", Price: 59.99},
	{ID: "6", Title: "The Rise And Fall Of Ziggy Stardust And The Spiders From Mars", Artist: "David Bowie", Price: 19.99},
	{ID: "7", Title: "Highway 61 Revisited", Artist: "Bob Dylan", Price: 9.99},
	{ID: "8", Title: "Untitled (Led Zeppelin IV)", Artist: "Led Zeppelin", Price: 29.99},
	{ID: "9", Title: "Exile On Main St.", Artist: "The Rolling Stones", Price: 109.99},
	{ID: "10", Title: "Funeral", Artist: "Arcade Fire", Price: 5.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, album := range albums {
		if album.ID == id {
			c.IndentedJSON(http.StatusOK, album)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}
