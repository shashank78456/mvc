package controller

import (
	"encoding/json"
	"net/http"

	"github.com/shashank78456/mvc/pkg/models"
	"github.com/shashank78456/mvc/pkg/types"
	"github.com/shashank78456/mvc/pkg/views"
)

func RenderSignup(writer http.ResponseWriter, request *http.Request) {
	t := views.SignupPage()
	t.Execute(writer, nil)
}

func HandleSignup(writer http.ResponseWriter, request *http.Request) {
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

	if !isUserExist {
		isUserTableNotEmpty, err := models.IsUserTableNotEmpty()
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in getting existing users"))
			return
		}

		userType := "client"
		if !isUserTableNotEmpty {
			userType = "superadmin"
		}

		hashedPassword, err := HashPassword(User.Password)

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error hashing password"))
			return
		}

		err = models.CreateUser(User.Username, hashedPassword, User.Name)
		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("User Creation Failed"))
			return
		}

		cookie := SendToken(writer, request, User.Username, userType)
		http.SetCookie(writer, &cookie)
		writer.Header().Set("Content-Type", "application/json")

		validity := make(map[string]interface{})
		validity["isValid"] = true
		validity["userType"] = userType
		response, err := json.Marshal(validity)

		if err != nil {
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

		if err != nil {
			writer.WriteHeader(http.StatusInternalServerError)
			writer.Write([]byte("Error in JSON Marshal"))
			return
		}

		writer.Write(response)
	}
}
