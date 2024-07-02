package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
)

func RenderAdmin(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	
	if(page=="home") {
		t := views.AdminHomePage()
		books, err := models.FetchBooks(0)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else if(page=="add") {
		t := views.AdminPromptPage()
		books, err := models.FetchBooks(0)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else if(page=="requests") {
		t := views.AdminRequestsPage()
		requests, err := models.FetchRequests()

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch requests"))
			return
		}

		Requests := types.RequestList{
			Requests: requests,
		}

		t.Execute(writer, Requests)

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func RenderSuperAdmin(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	
	if(page=="home") {
		t := views.SuperAdminHomePage()
		books, err := models.FetchBooks(0)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else if(page=="add") {
		t := views.SuperAdminPromptPage()
		books, err := models.FetchBooks(0)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		Books := types.BookList{
			Books: books,
		}

		t.Execute(writer, Books)

	} else if(page=="requests") {
		t := views.SuperAdminRequestsPage()
		requests, err := models.FetchRequests()

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch requests"))
			return
		}

		Requests := types.RequestList{
			Requests: requests,
		}

		t.Execute(writer, Requests)

	} else if(page=="adreq") {
		t := views.SuperAdminSuperPage()
		users, err := models.FetchUsersWithAdminRequest()

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not get Users with Admin Request"))
			return
		}

		Users := types.UserList{
			Users: users,
		}

		t.Execute(writer, Users)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func AddNewBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	isAdded, err := models.AddNewBook(Book.Bookname, Book.Author)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Add New Book"))
		return
	}

	status := make(map[string]bool)
	status["isAdded"] = isAdded
	response, err := json.Marshal(status)

	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error in JSON Marshal"))
		return
	}

	writer.Write(response)
}

func AddBook(writer http.ResponseWriter, request *http.Request) {
	var Book types.Book
	err := json.NewDecoder(request.Body).Decode(&Book)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.AddBook(Book.BookID)
	if (err!=nil) {
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
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.RemoveBook(Book.BookID)
	if (err!=nil) {
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
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	isBorrowed, err := models.IsBorrowed(Book.BookID)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Unable to fetch Requests"))
		return
	}

	if(isBorrowed) {
		status := make(map[string]bool)
		status["isDeleted"] = false
		response, err := json.Marshal(status)
	
		if (err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}
	
		writer.Write(response)

	} else {
		err =  models.DeleteBook(Book.BookID)
		if (err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Failed to Delete Book"))
			return
		}
	
		err = models.DeleteRequest(Book.BookID)
		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Failed to Delete Request"))
			return
	
		}
		status := make(map[string]bool)
		status["isDeleted"] = true
		response, err := json.Marshal(status)
	
		if (err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}
	
		writer.Write(response)
	}

}

func AcceptRequest(writer http.ResponseWriter, request *http.Request) {
	var Request types.Request
	err := json.NewDecoder(request.Body).Decode(&Request)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.AcceptRequest(Request.RequestID)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Accept Request"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Accepted Successfully"))	
}

func AcceptAdmin(writer http.ResponseWriter, request *http.Request) {
	var User types.User
	err := json.NewDecoder(request.Body).Decode(&User)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.HandleAdminRequest(User.UserID, true)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Accept Request"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Accepted Successfully"))
}

func DenyAdmin(writer http.ResponseWriter, request *http.Request) {
	var User types.User
	err := json.NewDecoder(request.Body).Decode(&User)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	err = models.HandleAdminRequest(User.UserID, false)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Failed to Accept Request"))
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Denied Successfully"))
}