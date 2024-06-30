package models

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/types"
)

func CreateRequest(userID int, bookID int) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return false, err
	}
	
	checksql := "SELECT * FROM Requests WHERE userID = ? AND bookID = ?"
	rows, err := db.Query(checksql, userID, bookID)
	if(err!=nil) {
		fmt.Println("Failed to fetch existing Requests")
		return false, err
	}

	if(!rows.Next()) {
		sql := "INSERT INTO Requests (userID, bookID) VALUES (?, ?)"
		_, err = db.Exec(sql, userID, bookID)

		if(err!=nil) {
			fmt.Println("Creating New Request Failed")
			return false, err
		}
		return true, nil

	} else {
		return false, nil
	}
}

func AcceptRequest(requestID int) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "UPDATE Requests SET isBorrowed = 1 AND isAccepted = 1 WHERE requestID = ?"
	_, err = db.Exec(sql, requestID)
	if(err!=nil) {
		fmt.Println("Updating Request Failed")
		return err
	}

	getsql := "SELECT bookid FROM Requests WHERE requestid = ?"

	rows, err := db.Query(getsql, requestID)
	if(err!=nil) {
		fmt.Println("Fetching BookID Failed")
		return err
	}

	var bookID int
	for rows.Next() {
		err := rows.Scan(&bookID)
		if(err!=nil) {
			fmt.Println("Error scanning rows")
			return err
		}
	}

	booksql := "UPDATE Books SET quantity = quantity - 1 WHERE bookID = ?"
	_, err = db.Exec(booksql, bookID)
	if(err!=nil) {
		fmt.Println("Updating Book Quantity Failed")
		return err
	}

	return nil
}

func CloseRequest(userID int, bookID int) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "UPDATE Requests SET isBorrowed = 0 WHERE userID = ? and bookID = ?"
	_, err = db.Exec(sql, userID, bookID)
	if(err!=nil) {
		fmt.Println("Closing Request Failed")
		return err
	}

	booksql := "UPDATE Books SET quantity = quantity + 1 WHERE bookID = ?"
	_, err = db.Exec(booksql, bookID)
	if(err!=nil) {
		fmt.Println("Updating Book Quantity Failed")
		return err
	}

	return nil
}

func DeleteRequest(bookID int) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "DELETE FROM Requests WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)
	if(err!=nil) {
		fmt.Println("Deleting Request Failed")
		return err
	}

	return nil
}

func FetchRequests() (types.RequestList, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return types.RequestList{}, err
	}

	sql := "SELECT * FROM Requests WHERE isAccepted = 0"
	rows, err := db.Query(sql)
	if(err!=nil) {
		fmt.Println("Failed to Fetch Requests")
		return types.RequestList{}, err
	}

	var fetchRequests []types.Request
	for rows.Next() {
		var request types.Request
		err := rows.Scan(&request)
		if(err==nil) {
			fmt.Println("Error in scanning rows")
			return types.RequestList{}, err
		}
		fetchRequests = append(fetchRequests, request)
	}

	var Requests types.RequestList
	Requests.Requests = fetchRequests
	return Requests, nil
}

func FetchBorrowedBooks(userID int) (types.BookList, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return types.BookList{}, err
	}

	sql := "SELECT bookid FROM Requests WHERE isAccepted = 1 AND isBorrowed = 1"
	rows, err := db.Query(sql)
	if(err!=nil) {
		fmt.Println("Failed to Fetch Borrowed Books")
		return types.BookList{}, err
	}

	var fetchBorrowedBooks []types.Book
	for rows.Next() {
		var bookid int
		err := rows.Scan(&bookid)
		if(err==nil) {
			fmt.Println("Error in scanning rows")
			return types.BookList{}, err
		}
		book, err := fetchBook(bookid)
		if(err!=nil) {
			fmt.Println("Error in fetching book")
			return types.BookList{}, err
		}
		fetchBorrowedBooks = append(fetchBorrowedBooks, book)
	}

	var Books types.BookList
	Books.Books = fetchBorrowedBooks
	return Books, nil
}

func FetchHistory(userID int) (types.BookList, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return types.BookList{}, err
	}

	sql := "SELECT bookid FROM Requests WHERE userID = ? AND isAccepted = 1"
	rows, err := db.Query(sql, userID)
	if(err!=nil) {
		fmt.Println("Failed to Fetch History")
		return types.BookList{}, err
	}

	var fetchHistory []types.Book
	for rows.Next() {
		var bookid int
		err := rows.Scan(&bookid)
		if(err==nil) {
			fmt.Println("Error in scanning rows")
			return types.BookList{}, err
		}
		book, err := fetchBook(bookid)
		if(err!=nil) {
			fmt.Println("Error in fetching book")
			return types.BookList{}, err
		}
		fetchHistory = append(fetchHistory, book)
	}

	var Books types.BookList
	Books.Books = fetchHistory
	return Books, nil
}

func IsBorrowed(bookID int) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return false, err
	}

	sql := "SELECT * FROM Requests WHERE bookID = ? AND isBorrowed = 1"
	rows, err := db.Query(sql, bookID)
	if(err!=nil) {
		fmt.Println("Failed to Fetch Requests")
		return false, err
	}

	return rows.Next(), nil
}