package models

import (
	"fmt"
	"github.com/shashank78456/mvc/pkg/types"
)

func AddNewBook(bookname string, author string) (bool, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return false, err
	}
	
	checksql := "SELECT * FROM Books WHERE bookname = ? AND author = ?"
	rows, err := db.Query(checksql, bookname, author)
	if(err!=nil) {
		fmt.Println("Failed to fetch existing books")
		return false, err
	}

	if(!rows.Next()) {
		sql := "INSERT INTO Books (bookname, author) VALUES (?, ?)"
		_, err = db.Exec(sql, bookname, author)

		if(err!=nil) {
			fmt.Println("Inserting New Book Failed")
			return false, err
		}
		return true, nil

	} else {
		return false, nil
	}
}

func AddBook(bookID int) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "UPDATE Books SET quantity = quantity + 1 WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)

	if(err!=nil) {
		fmt.Println("Adding Book Failed")
		return err
	}
	return nil
}

func RemoveBook(bookID int) error {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}

	sql := "UPDATE Books SET quantity = quantity - 1 WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)

	if(err!=nil) {
		fmt.Println("Removing Book Failed")
		return err
	}
	return nil
}

func DeleteBook(bookID int) (error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return err
	}
	
	sql := "DELETE FROM Books WHERE bookID = ?"
	_, err = db.Exec(sql, bookID)

	if(err!=nil) {
		fmt.Println("Deleting Book Failed")
		return err
	}
	return nil
}

func FetchBooks(minQuantity int) (types.BookList, error) {
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return types.BookList{}, err
	}

	sql := "SELECT * FROM Books WHERE quantity >= ?"
	rows, err := db.Query(sql, minQuantity)

	if(err!=nil) {
		fmt.Println("Fetching Books Failed")
		return types.BookList{}, err
	}

	var fetchBooks []types.Book
	for rows.Next() {
		var book types.Book
		err := rows.Scan(&book.Bookname)
		if(err!=nil) {
			fmt.Println("Error in scanning rows")
			return types.BookList{}, err
		}
		fetchBooks = append(fetchBooks, book)
	}

	var Books types.BookList
	Books.Books = fetchBooks
 	return Books, nil
}

func fetchBook(bookid int) (types.Book, error){
	db, err := Connection()
	if(err!=nil) {
		fmt.Println("Error in connecting to DB")
		return types.Book{}, err
	}

	sql := "SELECT * FROM Books WHERE bookid = ?"
	rows, err := db.Query(sql, bookid)
	if(err!=nil) {
		fmt.Println("Fetching Book Failed")
		return types.Book{}, err
	}

	var Book types.Book
	for rows.Next() {
		err := rows.Scan(&Book)
		if(err!=nil) {
			fmt.Println("Error in scanning rows")
			return types.Book{}, err
		}
	}

	return Book, nil
}