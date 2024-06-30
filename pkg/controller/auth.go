package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))
var expiry = time.Now().Add(time.Hour * 2)

func verifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if (err!=nil) {
		return "",err
	}

	if (!token.Valid) {
		return "", fmt.Errorf("Forbidden")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if (!ok) {
		return "", fmt.Errorf("username not found")
	}

	username, ok := claims["username"].(string)
	
	if (!ok) {
		return "", fmt.Errorf("username not found")
	}

	return username, nil
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"username": username,
		"expiry": expiry,
	})

	tokenString, err := token.SignedString(secretKey)
	if (err!=nil) {
		fmt.Println("Token Creation Failed", err)
		return "", err
	}

	return tokenString, nil
}

func Authenticator(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tokenString := request.Header.Get("Authorisation")[len("Bearer"):]
		if (tokenString == "") {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("No Token Found"))
			return
		}

		username, err := verifyToken(tokenString)

		if (err!=nil) {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Forbidden"))
			return
		}

		request = request.WithContext(context.WithValue(request.Context(), "username", username))
		next.ServeHTTP(writer, request)
	})
}

func SendToken(writer http.ResponseWriter, request *http.Request, username string) http.Cookie {
	token, err := createToken(username)

	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Token Creation failed"))
		return http.Cookie{}
	}
	
	cookie := http.Cookie{
		Name: "token",
		Value: token,
		Expires: expiry,
		Path: "/",
	}

	return cookie
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 20)
	if (err!=nil) {
		return "",err
	}

	return string(hashedPassword), nil
}

func IsPasswordValid(password string, originalPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(password))
	return err==nil
}
