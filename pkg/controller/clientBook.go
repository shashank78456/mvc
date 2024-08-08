package controller

import (
	"encoding/json"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"net/http"
)

func RequestBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	username := request.Context().Value("username").(string)
	userID, err := models.GetUserID(username)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get UserID"))
		return
	}

	hasAlreadyRequested, err := models.CreateRequest(userID, Book.BookID)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Creation of Request Failed "))
		return
	}

	status := make(map[string]bool)
	status["hasAlreadyRequested"] = hasAlreadyRequested
	response, err := json.Marshal(status)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error in JSON Marshal"))
		return
	}

	writer.Write(response)
}

func ReturnBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	username := request.Context().Value("username").(string)
	userID, err := models.GetUserID(username)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get UserID"))
		return
	}

	err = models.CloseRequest(userID, Book.BookID)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not return book"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Book Returned Successfully"))
}
