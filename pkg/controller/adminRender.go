package controller

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
)

func RenderAdmin(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	name, err := models.GetName(request.Context().Value("username").(string))

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not fetch name"))
		return
	}

	if page == "home" {
		t := views.AdminHomePage()
		books, err := models.FetchBooks(0, false, 0)

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

	} else if page == "add" {
		t := views.AdminPromptPage()
		books, err := models.FetchBooks(0, false, 0)

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

	} else if page == "requests" {
		t := views.AdminRequestsPage()
		requests, err := models.FetchRequests()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch requests"))
			return
		}

		Requests := types.RequestList{
			Requests: requests,
			Name:     name,
		}

		t.Execute(writer, Requests)

	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}

func RenderSuperAdmin(writer http.ResponseWriter, request *http.Request) {
	page := mux.Vars(request)["page"]
	name, err := models.GetName(request.Context().Value("username").(string))

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not fetch name"))
		return
	}

	if page == "home" {
		t := views.SuperAdminHomePage()
		books, err := models.FetchBooks(0, false, 0)

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

	} else if page == "add" {
		t := views.SuperAdminPromptPage()
		books, err := models.FetchBooks(0, false, 0)

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

	} else if page == "requests" {
		t := views.SuperAdminRequestsPage()
		requests, err := models.FetchRequests()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not fetch requests"))
			return
		}

		Requests := types.RequestList{
			Requests: requests,
			Name:     name,
		}

		t.Execute(writer, Requests)

	} else if page == "adminrequests" {
		t := views.SuperAdminSuperPage()
		users, err := models.FetchUsersWithAdminRequest()

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not get Users with Admin Request"))
			return
		}

		Users := types.UserList{
			Users: users,
			Name:  name,
		}

		t.Execute(writer, Users)
	} else {
		writer.WriteHeader(http.StatusNotFound)
	}
}
