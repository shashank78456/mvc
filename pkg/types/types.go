package types

//make chnages in type according to the json data
type User struct {
	UserID int
	Username string
	UserType string
	Name string
	Password string
}

type Book struct {
	BookID int
	Bookname string
	Author string
	Quantity int
}

type Request struct {
	RequestID int
	UserID int
	BookID int
	Username string
	Bookname string
	Author string
}

type BookList struct {
	Books []Book
}

type UserList struct {
	Users []User
}

type RequestList struct {
	Requests []Request
}

type BookData struct {
	IsSuperAdmin bool
	Books BookList
}

type UserData struct {
	IsSuperAdmin bool
	Users UserList
}

type RequestData struct {
	IsSuperAdmin bool
	Requests RequestList
}
