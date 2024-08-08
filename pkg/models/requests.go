package models

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/types"
)

func CreateRequest(userID int, bookID int) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}

	checksql := "SELECT * FROM Requests WHERE userID = ? AND bookID = ? AND ((isBorrowed = 0 AND isAccepted = 0) OR (isBorrowed = 1 AND isAccepted = 1))"
	rows, err := db.Query(checksql, userID, bookID)
	if err != nil {
		fmt.Println("Failed to fetch existing Requests", err)
		return false, err
	}

	if !rows.Next() {
		sql := "INSERT INTO Requests (userID, bookID) VALUES (?, ?)"
		_, err = db.Exec(sql, userID, bookID)

		if err != nil {
			fmt.Println("Creating New Request Failed", err)
			return false, err
		}
		return true, nil

	} else {
		return false, nil
	}
}

func AcceptRequest(requestID int) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}

	getsql := "SELECT bookid FROM Requests WHERE requestID = ?"

	rows, err := db.Query(getsql, requestID)
	if err != nil {
		fmt.Println("Fetching BookID Failed", err)
		return false, err
	}

	var bookID int
	for rows.Next() {
		err := rows.Scan(&bookID)
		if err != nil {
			fmt.Println("Error scanning rows", err)
			return false, err
		}
	}

	checksql := "SELECT quantity FROM Books WHERE bookID = ?"
	rows, err = db.Query(checksql, bookID)
	if err != nil {
		fmt.Println("Fetching Book Quantity Failed", err)
		return false, err
	}

	var quantity int
	for rows.Next() {
		err := rows.Scan(&quantity)
		if err != nil {
			fmt.Println("Error scanning rows", err)
			return false, err
		}
	}

	if quantity > 0 {

		sql := "UPDATE Requests SET isBorrowed = 1, isAccepted = 1 WHERE requestID = ?"
		_, err = db.Exec(sql, requestID)
		if err != nil {
			fmt.Println("Updating Request Failed", err)
			return false, err
		}

		booksql := "UPDATE Books SET quantity = quantity - 1 WHERE bookID = ?"
		_, err = db.Exec(booksql, bookID)
		if err != nil {
			fmt.Println("Updating Book Quantity Failed", err)
			return false, err
		}

		return true, nil
	} else {
		return false, nil
	}
}

func DenyRequest(requestID int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "DELETE FROM Requests WHERE requestID = ?"
	_, err = db.Exec(sql, requestID)
	if err != nil {
		fmt.Println("Deleting Request Failed", err)
		return err
	}

	return nil
}

func CloseRequest(userID int, bookID int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "UPDATE Requests SET isBorrowed = 0 WHERE userID = ? and bookID = ?"
	_, err = db.Exec(sql, userID, bookID)
	if err != nil {
		fmt.Println("Closing Request Failed", err)
		return err
	}

	booksql := "UPDATE Books SET quantity = quantity + 1 WHERE bookID = ?"
	_, err = db.Exec(booksql, bookID)
	if err != nil {
		fmt.Println("Updating Book Quantity Failed", err)
		return err
	}

	return nil
}

func DeleteRequest(bookID int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "DELETE FROM Requests WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)
	if err != nil {
		fmt.Println("Deleting Request Failed", err)
		return err
	}

	return nil
}

func FetchRequests() ([]types.Request, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return []types.Request{}, err
	}

	sql := "SELECT requestID, userID, bookID FROM Requests WHERE isAccepted = 0"
	rows, err := db.Query(sql)
	if err != nil {
		fmt.Println("Failed to Fetch Requests", err)
		return []types.Request{}, err
	}

	var fetchRequests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request.RequestID, &request.UserID, &request.BookID)
		if err != nil {
			fmt.Println("Error in scanning request rows", err)
			return []types.Request{}, err
		}

		booksql := "SELECT bookname, author FROM Books WHERE bookID = ?"
		bookrows, err := db.Query(booksql, request.BookID)
		if err != nil {
			fmt.Println("Failed to Fetch Books", err)
			return []types.Request{}, err
		}
		for bookrows.Next() {
			err := bookrows.Scan(&request.Bookname, &request.Author)
			if err != nil {
				fmt.Println("Error in scanning book rows", err)
				return []types.Request{}, err
			}
		}

		usersql := "SELECT username FROM Users WHERE userID = ?"
		userrows, err := db.Query(usersql, request.UserID)
		if err != nil {
			fmt.Println("Failed to Fetch Users", err)
			return []types.Request{}, err
		}
		for userrows.Next() {
			err := userrows.Scan(&request.Username)
			if err != nil {
				fmt.Println("Error in scanning user rows", err)
				return []types.Request{}, err
			}
		}

		fetchRequests = append(fetchRequests, request)
	}

	return fetchRequests, nil
}

func FetchBorrowedBooks(userID int) ([]types.Book, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return []types.Book{}, err
	}

	sql := "SELECT bookid FROM Requests WHERE userID = ? AND isAccepted = 1 AND isBorrowed = 1"
	rows, err := db.Query(sql, userID)
	if err != nil {
		fmt.Println("Failed to Fetch Borrowed Books", err)
		return []types.Book{}, err
	}

	var fetchBorrowedBooks []types.Book
	for rows.Next() {
		var bookid int
		err := rows.Scan(&bookid)
		if err != nil {
			fmt.Println("Error in scanning rows", err)
			return []types.Book{}, err
		}
		book, err := fetchBook(bookid)
		if err != nil {
			fmt.Println("Error in fetching book", err)
			return []types.Book{}, err
		}
		fetchBorrowedBooks = append(fetchBorrowedBooks, book)
	}

	return fetchBorrowedBooks, nil
}

func FetchHistory(userID int) ([]types.Book, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return []types.Book{}, err
	}

	sql := "SELECT bookid FROM Requests WHERE userID = ? AND isAccepted = 1"
	rows, err := db.Query(sql, userID)
	if err != nil {
		fmt.Println("Failed to Fetch History", err)
		return []types.Book{}, err
	}

	var fetchHistory []types.Book
	for rows.Next() {
		var bookid int
		err := rows.Scan(&bookid)
		if err != nil {
			fmt.Println("Error in scanning rows", err)
			return []types.Book{}, err
		}
		book, err := fetchBook(bookid)
		if err != nil {
			fmt.Println("Error in fetching book", err)
			return []types.Book{}, err
		}
		fetchHistory = append(fetchHistory, book)
	}

	return fetchHistory, nil
}

func IsBorrowed(bookID int) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}

	sql := "SELECT * FROM Requests WHERE bookID = ? AND isBorrowed = 1"
	rows, err := db.Query(sql, bookID)
	if err != nil {
		fmt.Println("Failed to Fetch Requests", err)
		return false, err
	}

	return rows.Next(), nil
}

func DeleteRequestsOfUser(userid int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "DELETE FROM Requests WHERE userid = ?"
	_, err = db.Exec(sql, userid)
	if err != nil {
		fmt.Println("Failed to Delete Requests", err)
		return err
	}

	return nil
}

func IsAlreadyRequestedOrBorrowed(userID int, bookID int) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}

	checksql := "SELECT * FROM Requests WHERE userID = ? AND bookID = ? AND ((isBorrowed = 0 AND isAccepted = 0) OR (isBorrowed = 1 AND isAccepted = 1))"
	rows, err := db.Query(checksql, userID, bookID)
	if err != nil {
		fmt.Println("Failed to fetch existing Requests", err)
		return false, err
	}

	if !rows.Next() {
		return false, nil
	} else {
		return true, nil
	}
}
