package controller

import (
	"github.com/gorilla/mux"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
	"net/http"
)

func RenderClient(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	name, err := models.GetName(request.Context().Value("username").(string))

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not fetch name"))
		return
	}

	if page == "home" {
		t := views.ClientHomePage()

		userID, err := models.GetUserID(name)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not get userID"))
			return
		}

		books, err := models.FetchBooks(1, true, userID)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		Books := types.BookList{
			Books: books,
			Name:  name,
		}

		t.Execute(writer, Books)

	} else if page == "history" {
		username := request.Context().Value("username").(string)
		userID, err := models.GetUserID(username)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch UserID"))
			return
		}

		t := views.ClientHistoryPage()
		books, err := models.FetchHistory(userID)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch history"))
			return
		}

		Books := types.BookList{
			Books: books,
			Name:  name,
		}

		t.Execute(writer, Books)

	} else if page == "return" {
		username := request.Context().Value("username").(string)
		userID, err := models.GetUserID(username)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch UserID"))
			return
		}

		t := views.ClientReturnPage()
		books, err := models.FetchBorrowedBooks(userID)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch borrowed books"))
			return
		}

		Books := types.BookList{
			Books: books,
			Name:  name,
		}

		t.Execute(writer, Books)

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}
