package models

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/types"
)

func AddNewBook(bookname string, author string, quantity int) (bool, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return false, err
	}

	checksql := "SELECT * FROM Books WHERE bookname = ? AND author = ?"
	rows, err := db.Query(checksql, bookname, author)
	if err != nil {
		fmt.Println("Failed to fetch existing books", err)
		return false, err
	}

	if !rows.Next() {
		sql := "INSERT INTO Books (bookname, author, quantity) VALUES (?, ?, ?)"
		_, err = db.Exec(sql, bookname, author, quantity)

		if err != nil {
			fmt.Println("Inserting New Book Failed", err)
			return false, err
		}
		return true, nil

	} else {
		return false, nil
	}
}

func AddBook(bookID int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "UPDATE Books SET quantity = quantity + 1 WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)

	if err != nil {
		fmt.Println("Adding Book Failed", err)
		return err
	}
	return nil
}

func RemoveBook(bookID int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "UPDATE Books SET quantity = quantity - 1 WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)

	if err != nil {
		fmt.Println("Removing Book Failed", err)
		return err
	}
	return nil
}

func DeleteBook(bookID int) error {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return err
	}

	sql := "DELETE FROM Books WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)

	if err != nil {
		fmt.Println("Deleting Book Failed", err)
		return err
	}
	return nil
}

func FetchBooks(minQuantity int, check bool, userID int) ([]types.Book, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return []types.Book{}, err
	}

	sql := "SELECT * FROM Books WHERE quantity >= ?"
	rows, err := db.Query(sql, minQuantity)

	if err != nil {
		fmt.Println("Fetching Books Failed", err)
		return []types.Book{}, err
	}

	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.BookID, &book.Bookname, &book.Author, &book.Quantity)
		if err != nil {
			fmt.Println("Error in scanning rows", err)
			return []types.Book{}, err
		}
		fetchBooks = append(fetchBooks, book)
	}

	if check {
		for i := 0; i < len(fetchBooks); i++ {
			fetchBooks[i].IsAlreadyRequestedOrBorrowed, err = IsAlreadyRequestedOrBorrowed(userID, fetchBooks[i].BookID)
			if err != nil {
				fmt.Println("Error in getting borrowed status", err)
				return []types.Book{}, err
			}
		}
	}

	return fetchBooks, nil
}

func fetchBook(bookid int) (types.Book, error) {
	db, err := Connection()
	if err != nil {
		fmt.Println("Error in connecting to DB", err)
		return types.Book{}, err
	}

	sql := "SELECT * FROM Books WHERE bookid = ?"
	rows, err := db.Query(sql, bookid)
	if err != nil {
		fmt.Println("Fetching Book Failed", err)
		return types.Book{}, err
	}

	var Book types.Book
	for rows.Next() {
		err := rows.Scan(&Book.BookID, &Book.Bookname, &Book.Author, &Book.Quantity)
		if err != nil {
			fmt.Println("Error in scanning rows", err)
			return types.Book{}, err
		}
	}

	return Book, nil
}
