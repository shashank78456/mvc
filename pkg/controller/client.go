package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
)

func RenderClient(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	if(page=="home") {
		t := views.ClientHomePage()
		books, err := models.FetchBooks(1)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else if(page=="history") {
		username := request.Context().Value("username").(string)
		userID, err := models.GetUserID(username)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch UserID"))
			return
		}

		t := views.ClientHistoryPage()
		books, err := models.FetchHistory(userID)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch history"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else if(page=="return") {
		username := request.Context().Value("username").(string)
		userID, err := models.GetUserID(username)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch UserID"))
			return
		}

		t := views.ClientReturnPage()
		books, err := models.FetchBorrowedBooks(userID)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch borrowed books"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func RequestBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	username := request.Context().Value("username").(string)
	userID, err := models.GetUserID(username)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get UserID"))
		return
	}
	
	hasAlreadyRequested, err := models.CreateRequest(userID, Book.BookID)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Creation of Request Failed "))
		return
	}

	status := make(map[string]bool)
	status["hasAlreadyRequested"] = hasAlreadyRequested
	response, err := json.Marshal(status)

	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error in JSON Marshal"))
		return
	}

	writer.Write(response)	
}

func ReturnBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	username := request.Context().Value("username").(string)
	userID, err := models.GetUserID(username)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get UserID"))
		return
	}

	err = models.CloseRequest(userID, Book.BookID)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not return book"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Book Returned Successfully"))
}

func AdminRequest(writer http.ResponseWriter, request *http.Request) {
	username := request.Context().Value("username").(string)
	userID, err := models.GetUserID(username)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get UserID"))
		return
	}

	hasAlreadyRequested, err := models.HasAlreadyRequested(userID)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not check if already requested"))
		return
	}

	if(!hasAlreadyRequested) {
		err := models.RequestForAdmin(userID)
		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not request for admin"))
			return
		}
		
	}

	status := make(map[string]bool)
	status["hasAlreadyRequested"] = hasAlreadyRequested
	response, err := json.Marshal(status)

	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error in JSON Marshal"))
		return
	}

	writer.Write(response)
}