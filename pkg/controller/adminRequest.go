package controller

import (
	"encoding/json"
	"net/http"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
)

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