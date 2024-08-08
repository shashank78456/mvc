package controller

import (
	"encoding/json"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"net/http"
)

func AddNewBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	isAdded, err := models.AddNewBook(Book.Bookname, Book.Author, Book.Quantity)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Add New Book"))
		return
	}

	status := make(map[string]bool)
	status["isAdded"] = isAdded
	response, err := json.Marshal(status)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error in JSON Marshal"))
		return
	}

	writer.Write(response)
}

func AddBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.AddBook(Book.BookID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Add Book"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Added Successfully"))
}

func RemoveBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.RemoveBook(Book.BookID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Remove Book"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Removed Successfully"))
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	isBorrowed, err := models.IsBorrowed(Book.BookID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Unable to fetch Requests"))
		return
	}

	if isBorrowed {
		status := make(map[string]bool)
		status["isDeleted"] = false
		response, err := json.Marshal(status)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)

	} else {
		err = models.DeleteBook(Book.BookID)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Failed to Delete Book"))
			return
		}

		err = models.DeleteRequest(Book.BookID)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Failed to Delete Request"))
			return

		}
		status := make(map[string]bool)
		status["isDeleted"] = true
		response, err := json.Marshal(status)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)
	}

}
