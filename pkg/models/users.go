package models

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/types"
)

func CreateUser(username string, password string, name string) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "INSERT INTO Users (username, password, name) VALUES (?, ?, ?)"
	_, err = db.Exec(sql, username, password, name)

	if(err!=nil) {
		fmt.Println("Creating New User Failed")
		return err
	}
	return nil
}

func RequestForAdmin(userID int) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return false, err
	}

	sql := "SELECT hasAdminRequest FROM Users WHERE userID = ?"
	rows, err := db.Query(sql, userID)
	if(err!=nil) {
		fmt.Println("Failed to fetch Status")
		return false, err
	}

	var hasAdminRequest int
	for rows.Next() {
		err := rows.Scan(&hasAdminRequest)
		if(err!=nil) {
			fmt.Println("Error in scanning rows")
			return false, err
		}
	}

	return hasAdminRequest==1, nil
}

func HandleAdminRequest(userID int, isAccepted bool) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	status := 0
	if(isAccepted) {
		status = 1
	}
	sql := "UPDATE Users SET isAccepted = ? WHERE userID = ?"
	_, err = db.Exec(sql, status, userID)
	if(err!=nil) {
		fmt.Println("Failed to update status")
		return err
	}
	return nil
}

func FetchUsersWithAdminRequest() (types.UserList, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return types.UserList{}, err
	}

	sql := "SELECT * FROM Users WHERE hasAdminRequest = 1"
	rows, err := db.Query(sql)
	if(err!=nil) {
		fmt.Println("Failed to fetch Users")
		return types.UserList{}, err
	}

	var fetchUsers []types.User
	for rows.Next() {
		var user types.User
		err := rows.Scan(&user)
		if(err!=nil) {
			fmt.Println("Error scanning rows")
			return types.UserList{}, err
		}
	}

	var Users types.UserList
	Users.Users = fetchUsers
	return Users, nil
}

func IsSuperAdmin(userID int) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return false, err
	}

	sql := "SELECT isSuperAdmin FROM Users WHERE userID = ?"
	rows, err := db.Query(sql, userID)
	if(err!=nil) {
		fmt.Println("Failed to fetch Adminlevel")
		return false, err
	}

	var isSuperAdmin bool
	for rows.Next() {
		err := rows.Scan(&isSuperAdmin)
		if(err!=nil) {
			fmt.Println("Error scanning rows")
			return false, err
		}
	}
	return isSuperAdmin, nil
}

func GetUserID(username string) (int, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return -1, err
	}

	sql := "SELECT userID FROM Users WHERE username = ?"
	rows, err := db.Query(sql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch userID")
		return -1, err
	}

	var userID int
	for rows.Next() {
		err := rows.Scan(&userID)
		if(err!=nil) {
			fmt.Println("Error scanning rows")
			return -1, err
		}
	}
	return userID, nil
}

func GetPassword(username string) (string, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return "", err
	}

	sql := "SELECT password FROM Users WHERE username = ?"
	rows, err := db.Query(sql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch password")
		return "", err
	}

	var password string
	for rows.Next() {
		err := rows.Scan(&password)
		if(err!=nil) {
			fmt.Println("Error scanning rows")
			return "", err
		}
	}
	return password, nil
}

func GetUserType(username string) (string, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return "", err
	}

	sql := "SELECT userType FROM Users WHERE username = ?"
	rows, err := db.Query(sql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch userType")
		return "", err
	}

	var userType string
	for rows.Next() {
		err := rows.Scan(&userType)
		if(err!=nil) {
			fmt.Println("Error scanning rows")
			return "", err
		}
	}
	return userType, nil
}

func IsUserExist(username string) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return false, err
	}
	
	checksql := "SELECT * FROM Users WHERE username = ?"
	rows, err := db.Query(checksql, username)
	if(err!=nil) {
		fmt.Println("Failed to fetch existing users")
		return false, err
	}

	return rows.Next(), nil
}
