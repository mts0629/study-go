package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Data about a record album
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// Seed record album data
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughen and Clifford Brown", Artist: "Sarah Vaughen", Price: 39.99},
}

// Repond with the list of all albums as JSON
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// Add an album from JSON received in the request body
func postAlbums(c *gin.Context) {
	var newAlbum album

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	// Add the new album to the slice
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

// Locate the album whose ID value matches the id parameter sent by the client,
// then returns that album as a response
func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	// Loop over the list of albums, looking for an album
	// whose ID value matches the parameter
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
	router := gin.Default()

	// API endpoint: "/albums"
	router.GET("/albums", getAlbums)        // Get all albums by GET method
	router.GET("/albums/:id", getAlbumByID) // Get an album by GET method with id
	router.POST("/albums", postAlbums)      // Add an album by POST method

	// Launch a Gin router
	router.Run("localhost:8080")
}
