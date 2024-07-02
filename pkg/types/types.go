package types

type User struct {
	UserID int `json:"userid"`
	Username string `json:"username"`
	UserType string	`json:"userType"`
	Name string	`json:"name"`
	Password string	`json:"password"`
}

type Book struct {
	BookID int `json:"bookid"`
	Bookname string `json:"bookname"`
	Author string `json:"author"`
	Quantity int `json:"quantity"`
}

type Request struct {
	RequestID int `json:"requestID"`
	UserID int `json:"userid"`
	BookID int `json:"bookid"`
	Username string `json:"username"`
	Bookname string	`json:"bookname"`
	Author string `json:"author"`
}

type BookList struct {
	Books []Book `json:"books"`
}

type UserList struct {
	Users []User `json:"users"`
}

type RequestList struct {
	Requests []Request `json:"requests"`
}

