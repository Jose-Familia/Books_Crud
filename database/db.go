package database

import (
	"Books_Crud/models" // Importamos el modelo Book
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"os"
)

var DB *gorm.DB

func Init() {
	var err error

	dsn := os.Getenv("DATABASE_URL")

	DB, err = gorm.Open("postgres", dsn)

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Hacer la migración automática de la tabla Book
	DB.AutoMigrate(&models.Book{})

	// Llamamos a la función para insertar los 10 libros
	insertBooks()
}

func insertBooks() {
	books := []models.Book{
		{Title: "The Catcher in the Rye", Author: "J.D. Salinger", Year: 1951},
		{Title: "To Kill a Mockingbird", Author: "Harper Lee", Year: 1960},
		{Title: "1984", Author: "George Orwell", Year: 1949},
		{Title: "The Great Gatsby", Author: "F. Scott Fitzgerald", Year: 1925},
		{Title: "Moby-Dick", Author: "Herman Melville", Year: 1851},
		{Title: "War and Peace", Author: "Leo Tolstoy", Year: 1869},
		{Title: "Pride and Prejudice", Author: "Jane Austen", Year: 1813},
		{Title: "The Lord of the Rings", Author: "J.R.R. Tolkien", Year: 1954},
		{Title: "The Hobbit", Author: "J.R.R. Tolkien", Year: 1937},
		{Title: "Brave New World", Author: "Aldous Huxley", Year: 1932},
	}

	for _, book := range books {
		if err := DB.FirstOrCreate(&book, models.Book{Title: book.Title}).Error; err != nil {
			log.Printf("Failed to insert book: %v", err)
		}
	}

	log.Println("10 books have been registered in the database!")
}

func GetDB() *gorm.DB {
	return DB
}
