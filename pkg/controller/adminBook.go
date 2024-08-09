package controller

import (
	"encoding/json"
	"net/http"

	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
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

func EditBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.EditBook(Book.BookID, Book.Quantity)
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
		err = models.DeleteRequest(Book.BookID)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Failed to Delete Request"))
			return
		}

		err = models.DeleteBook(Book.BookID)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Failed to Delete Book"))
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
