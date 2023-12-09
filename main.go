package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, books)
	})
	r.Run(":8080")
}

type Books struct {
	ID     string "json:id"
	Title  string "json:title"
	Author string "json:author"
}

var books = []Books{
	{ID: "1", Title: "Go Rest API Basu", Author: "Basu"},
	{ID: "2", Title: "Harry Potter", Author: "J K Rowling"},
	{ID: "3", Title: "Something Story", Author: "Basu"},
	{ID: "4", Title: "Neko Cafe", Author: "Meow"},
	{ID: "5", Title: "Website Neko", Author: "Bammcool"},
}
