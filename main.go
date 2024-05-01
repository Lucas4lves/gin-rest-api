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


func GetAllAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK, albums)	
}

func main() {
	router := gin.Default()
	router.GET("/albums", GetAllAlbums)

	router.Run("localhost:8081")
}
