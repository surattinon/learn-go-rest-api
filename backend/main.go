// package name main
package main

// Import the Gin library and net/http
import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Create a new gin router
	r := gin.New()
	r.Use(cors.Default())
	// Create a new route with GET method
	r.GET("/books", func(c *gin.Context) {
		// Return a JSON response with status code 200
		c.JSON(http.StatusOK, books)
	})
	// Run the server on port 8080
	r.Run(":3000")
}

// Create a new struct for books
type Books struct {
	ID     string "json:id"
	Title  string "json:title"
	Author string "json:author"
}

// Create a new slice of books
var books = []Books{
	{ID: "1", Title: "Go Rest API Basu", Author: "Basu"},
	{ID: "2", Title: "Harry Potter", Author: "J K Rowling"},
	{ID: "3", Title: "Something Story", Author: "Basu"},
	{ID: "4", Title: "Neko Cafe", Author: "Meow"},
	{ID: "5", Title: "Website Neko", Author: "Bammcool"},
}
