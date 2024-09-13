package main

import (
	"log"
	"net/http"

	"Books_Crud/database"
	"Books_Crud/handlers"
	"github.com/julienschmidt/httprouter"
)

func main() {
	// Initialize database
	database.Init()

	router := httprouter.New()

	// Book routes
	router.POST("/books", handlers.CreateBook)
	router.GET("/books", handlers.GetBooks)
	router.GET("/books/:id", handlers.GetBook)
	router.PUT("/books/:id", handlers.UpdateBook)
	router.DELETE("/books/:id", handlers.DeleteBook)

	// Start server
	log.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
