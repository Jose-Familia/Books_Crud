package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"Books_Crud/database"
	"Books_Crud/models"
	"github.com/julienschmidt/httprouter"
)

// CreateBook - Handler to add a new book
func CreateBook(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var book models.Book
	json.NewDecoder(r.Body).Decode(&book)

	db := database.GetDB()
	if err := db.Create(&book).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// GetBooks - Handler to retrieve all books
func GetBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var books []models.Book

	db := database.GetDB()
	if err := db.Find(&books).Error; err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(books)
}

// GetBook - Handler to retrieve a book by ID
func GetBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))

	var book models.Book
	db := database.GetDB()
	if err := db.First(&book, id).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(book)
}

// UpdateBook - Handler to update a book by ID
func UpdateBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	var book models.Book
	db := database.GetDB()

	if db.First(&book, id).RecordNotFound() {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	json.NewDecoder(r.Body).Decode(&book)
	db.Save(&book)

	json.NewEncoder(w).Encode(book)
}

// DeleteBook - Handler to delete a book by ID
func DeleteBook(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("id"))
	var book models.Book

	db := database.GetDB()
	if db.First(&book, id).RecordNotFound() {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	db.Delete(&book)
	w.WriteHeader(http.StatusNoContent)
}
