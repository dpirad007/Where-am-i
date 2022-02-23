package main

import (
	"net/http"
    "github.com/dpirad007/Where-am-i/modals/student.go"
	"github.com/gin-gonic/gin"
)

// album represents data about a record album.
type album struct {
	ID     string  
	Title  string  
	Artist string 
	Price  float32 
}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// var students = []Student {
    // { ID: "1", name:"dion", rollNo:"501854", phoneNo:[]string{""}, email:[]string{""}, parentPhoneNo:[]string{""}},
// }

// get albums
func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

// post albums
func postAlbums(c *gin.Context) {
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)

}

func getStudent(c* gin.Context) {
    
}

func main() {
	router := gin.Default()
    router.GET("/student/:rollNo")
    router.GET("/teacher/:teacherId")

	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)

	router.Run("localhost:8080")
}
