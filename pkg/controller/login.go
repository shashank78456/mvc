package controller

import (
	"encoding/json"
	"net/http"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
)

func RenderLogin(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	t.Execute(writer, nil)
}

func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	var User types.User
	err := json.NewDecoder(request.Body).Decode(&User)
	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	isUserExist, err := models.IsUserExist(User.Username)
	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not check User existence"))
		return
	}

	originalPassword, err := models.GetPassword(User.Username)

	if(err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get Password"))
		return
	}

	if (isUserExist && IsPasswordValid(User.Password, originalPassword)) {
			cookie := SendToken(writer, request, User.Username)
			http.SetCookie(writer, &cookie)
			writer.Header().Set("Content-Type", "application/json")

			userType, err := models.GetUserType((User.Username))

			if(err!=nil) {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte("Could not get UserType"))
				return
			}

			validity := make(map[string]interface{})
			validity["isValid"] = true
			validity["userType"] = userType
			response, err := json.Marshal(validity)

			if (err!=nil) {
				writer.WriteHeader(http.StatusInternalServerError)
				writer.Write([]byte("Error in JSON Marshal"))
				return
			}

			writer.Write(response)

	} else {
		validity := make(map[string]interface{})
		validity["isValid"] = false
		validity["userType"] = ""
		response, err := json.Marshal(validity)

		if (err!=nil) {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)
	}
}