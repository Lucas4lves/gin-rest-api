package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//type with json annotations
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

//Mocked up data
var albums = []Album{
	{ID: "1", Title: "Rush of Blood To the Head", Artist:"Coldplay", Price: 17.99 },
	{ID: "2", Title: "X and Y", Artist:"Coldplay", Price: 45.89 },
	{ID: "3", Title: "Viva La Vida", Artist: "Coldplay", Price: 27.99 },
}

func PostAlbum(c *gin.Context){
	var newAlbum Album

	if err := c.BindJSON(&newAlbum); err != nil{
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func GetAlbumById(c *gin.Context){
	id := c.Param("id")	
	for _, a := range albums{
		if a.ID == id{
			c.IndentedJSON(http.StatusOK, a)
			return 
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"msg": "album not found"})
}

func GetAllAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)	
}

func main() {
	router := gin.Default()
	router.GET("/albums", GetAllAlbums)
	router.POST("/albums", PostAlbum)
	router.GET("/albums/:id", GetAlbumById)

	router.Run("localhost:8081")
}
