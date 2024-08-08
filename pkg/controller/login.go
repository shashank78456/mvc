package controller

import (
	"encoding/json"
	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
	"net/http"
)

func RenderLogin(writer http.ResponseWriter, request *http.Request) {
	t := views.LoginPage()
	t.Execute(writer, nil)
}

func HandleLogin(writer http.ResponseWriter, request *http.Request) {
	var User types.User
	err := json.NewDecoder(request.Body).Decode(&User)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Error Decoding struct"))
		return
	}

	isUserExist, err := models.IsUserExist(User.Username)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not check User existence"))
		return
	}

	originalPassword, err := models.GetPassword(User.Username)

	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Could not get Password"))
		return
	}

	if isUserExist && IsPasswordValid(User.Password, originalPassword) {
		User.UserType, err = models.GetUserType(User.Username)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not get usertype"))
			return
		}

		cookie := SendToken(writer, request, User.Username, User.UserType)
		http.SetCookie(writer, &cookie)
		writer.Header().Set("Content-Type", "application/json")

		userType, err := models.GetUserType((User.Username))

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Could not get UserType"))
			return
		}

		validity := make(map[string]interface{})
		validity["isExist"] = true
		validity["isValid"] = true
		validity["userType"] = userType
		response, err := json.Marshal(validity)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)

	} else if !isUserExist {
		validity := make(map[string]interface{})
		validity["isExist"] = false
		validity["isValid"] = false
		validity["userType"] = ""
		response, err := json.Marshal(validity)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)
	} else {
		validity := make(map[string]interface{})
		validity["isExist"] = true
		validity["isValid"] = false
		validity["userType"] = ""
		response, err := json.Marshal(validity)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)
	}
}
