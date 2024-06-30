package controller

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/views"
	"github.com/shashank78456/mvc/pkg/types"
)

func RenderAdmin(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	username := request.Context().Value("username").(string)
	userID, err := models.GetUserID(username)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get UserID"))
		return
	}
	
	IsSuperAdmin, err := models.IsSuperAdmin(userID)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get Admin Level"))
		return
	}
	
	if(page=="home") {
		t := views.AdminHomePage()
		Books, err := models.FetchBooks(0)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		BookData := types.BookData{
			Books: Books,
			IsSuperAdmin: IsSuperAdmin,
		}
		t.Execute(writer, BookData)

	} else if(page=="add") {
		t := views.AdminPromptPage()
		Books, err := models.FetchBooks(0)

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch books"))
			return
		}

		BookData := types.BookData{
			Books: Books,
			IsSuperAdmin: IsSuperAdmin,
		}
		t.Execute(writer, BookData)

	} else if(page=="requests") {
		t := views.AdminRequestsPage()
		Requests, err := models.FetchRequests()

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch requests"))
			return
		}

		RequestData := types.RequestData{
			Requests: Requests,
			IsSuperAdmin: IsSuperAdmin,
		}
		t.Execute(writer, RequestData)

	} else if(page=="adreq") {
		t := views.AdminSuperPage()
		Users, err := models.FetchUsersWithAdminRequest()

		if(err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not get Users with Admin Request"))
			return
		}

		UserData := types.UserData{
			Users: Users,
			IsSuperAdmin: IsSuperAdmin,
		}
		t.Execute(writer, UserData)
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