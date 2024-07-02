package models

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/types"
)

func CreateUser(username string, password string, name string) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return err
	}
	userType := "client"

	checksql := "SELECT * FROM USERS"
	rows, err := db.Query(checksql)

	if(err!=nil) {
		fmt.Println("Fetching Users Failed", err)
		return err
	}

	if(!rows.Next()) {
		userType = "superadmin"
	}

	sql := "INSERT INTO Users (username, password, name, userType) VALUES (?, ?, ?, ?)"
	_, err = db.Exec(sql, username, password, name, userType)

	if(err!=nil) {
		fmt.Println("Creating New User Failed", err)
		return err
	}
	return nil
}

func RequestForAdmin(userID int) (error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "UPDATE Users SET hasAdminRequest = 1 WHERE userID = ?"
	_, err = db.Exec(sql, userID)
	if(err!=nil) {
		fmt.Println("Failed to fetch Status", err)
		return err
	}

	return nil
}

func HandleAdminRequest(userID int, isAccepted bool) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	status := 0
	if(isAccepted) {
		status = 1
	}
	sql := "UPDATE Users SET isAccepted = ? WHERE userID = ?"
	_, err = db.Exec(sql, status, userID)
	if(err!=nil) {
		fmt.Println("Failed to update status", err)
		return err
	}
	return nil
}

func FetchUsersWithAdminRequest() ([]types.User, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return []types.User{}, err
	}

	sql := "SELECT userid, username FROM Users WHERE hasAdminRequest = 1"
	rows, err := db.Query(sql)
	if(err!=nil) {
		fmt.Println("Failed to fetch Users", err)
		return []types.User{}, err
	}

	var fetchUsers []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user.UserID, &user.Username)
		if(err!=nil) {
			fmt.Println("Error scanning rows", err)
			return []types.User{}, err
		}
	}

	return fetchUsers, nil
}

func GetUserID(username string) (int, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return -1, err
	}

	sql := "SELECT userID FROM Users WHERE username = ?"
	rows, err := db.Query(sql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch userID", err)
		return -1, err
	}

	var userID int
	for rows.Next() {
		err := rows.Scan(&userID)
		if(err!=nil) {
			fmt.Println("Error scanning rows", err)
			return -1, err
		}
	}
	return userID, nil
}

func GetPassword(username string) (string, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return "", err
	}

	sql := "SELECT password FROM Users WHERE username = ?"
	rows, err := db.Query(sql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch password", err)
		return "", err
	}

	var password string
	for rows.Next() {
		err := rows.Scan(&password)
		if(err!=nil) {
			fmt.Println("Error scanning rows", err)
			return "", err
		}
	}
	return password, nil
}

func GetUserType(username string) (string, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return "", err
	}

	sql := "SELECT userType FROM Users WHERE username = ?"
	rows, err := db.Query(sql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch userType", err)
		return "", err
	}

	var userType string
	for rows.Next() {
		err := rows.Scan(&userType)
		if(err!=nil) {
			fmt.Println("Error scanning rows", err)
			return "", err
		}
	}
	return userType, nil
}

func IsUserExist(username string) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}
	
	checksql := "SELECT * FROM Users WHERE username = ?"
	rows, err := db.Query(checksql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch existing users", err)
		return false, err
	}

	return rows.Next(), nil
}

func HasAlreadyRequested(userID int) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}

	sql := "SELECT hasAdminRequest FROM Users WHERE userID = ?"
	rows, err := db.Query(sql, userID)
	if(err!=nil) {
		fmt.Println("Failed to fetch Status", err)
		return false, err
	}
	
	var hasAdminRequest int
	for rows.Next() {
		err := rows.Scan(&hasAdminRequest)
		if(err!=nil) {
			fmt.Println("Error in scanning rows", err)
			return false, err
		}
	}

	return hasAdminRequest==1, nil

}