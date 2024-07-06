package controller

import (
	"encoding/json"
	"net/http"
	"github.com/shashank78456/mvc/pkg/models"
)

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