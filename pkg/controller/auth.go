package controller

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))
var expiry = time.Now().Add(time.Hour * 2)

func verifyToken(tokenString string) (string, string, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if (err!=nil) {
		return "", "", err
	}

	if (!token.Valid) {
		return "", "", fmt.Errorf("forbidden")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if (!ok) {
		return "", "", fmt.Errorf("username not found")
	}

	username, ok := claims["username"].(string)

	if (!ok) {
		return "", "", fmt.Errorf("usertype not found")
	}

	userType, ok := claims["userType"].(string)
	
	if (!ok) {
		return "", "", fmt.Errorf("username not found")
	}

	return username, userType, nil
}

func createToken(username string, userType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"userType" : userType,
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
		if(len(request.Header.Get("Cookie"))==0) {
			next.ServeHTTP(writer, request)
			return
		}
		var tokenName string
		var tokenString string
		cookieheader := request.Header.Get("Cookie")
		if(strings.Contains(cookieheader, "; ")) {
			cookies := strings.Split(cookieheader, "; ")
			for i := 0; i < len(cookies); i++ {
				cookie := strings.Split(cookies[i], "=")
				tokenName = cookie[0]
				if cookie[0]=="accesstoken" {
					tokenString = cookie[1]
					break
				}
			}
		} else {
			cookie := strings.Split(request.Header.Get("Cookie"), "=")
			tokenName = cookie[0]
			tokenString = cookie[1]
		}


		req := strings.Split(request.URL.String(), "/")[1]
		if(tokenName!="accesstoken") {
			if(req=="" || req=="signup") {
				ckie := http.Cookie{
					Name: tokenName,
					Value: "",
					Expires: time.Unix(0, 0),
					MaxAge: -1,
				}
				http.SetCookie(writer, &ckie)
				next.ServeHTTP(writer, request)
				return
			} else {
				writer.WriteHeader(http.StatusUnauthorized)
				writer.Write([]byte("No Token Found"))
				return
			}
		}
	
		if (tokenString == "") {
			if(req=="" || req=="signup") {
				ckie := http.Cookie{
					Name: "accesstoken",
					Value: tokenName,
					Expires: time.Unix(0, 0),
					MaxAge: -1,
				}
				http.SetCookie(writer, &ckie)
				next.ServeHTTP(writer, request)
				return
			}
			writer.WriteHeader(http.StatusOK)
			writer.Write([]byte("No Token Found"))
			return
		}

		username, userType, err := verifyToken(tokenString)

		if (err!=nil) {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Forbidden"))
			return
		}

		contxt := context.WithValue(request.Context(), "username", username)
		request = request.WithContext(contxt)

		if(req=="" || req=="signup") {
			http.Redirect(writer, request, fmt.Sprintf(`http://localhost:3000/%s/home`, userType), http.StatusFound)
			return
		} else if (!(req==userType)) {
			writer.WriteHeader(http.StatusUnauthorized)
			writer.Write([]byte("Forbidden"))
			return
		}

		next.ServeHTTP(writer, request)
	})
}

func SendToken(writer http.ResponseWriter, request *http.Request, username string, userType string) http.Cookie {
	token, err := createToken(username, userType)

	if (err!=nil) {
		writer.WriteHeader(http.StatusInternalServerError)
		writer.Write([]byte("Token Creation failed"))
		return http.Cookie{}
	}
	
	cookie := http.Cookie{
		Name: "accesstoken",
		Value: token,
		Expires: expiry,
		Path: "/",
	}

	return cookie
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if (err!=nil) {
		return "",err
	}

	return string(hashedPassword), nil
}

func IsPasswordValid(password string, originalPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(originalPassword), []byte(password))
	return err==nil
}
