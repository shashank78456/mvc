package types

type User struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	UserType string `json:"userType"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Book struct {
	BookID                       int    `json:"bookid"`
	Bookname                     string `json:"bookname"`
	Author                       string `json:"author"`
	Quantity                     int    `json:"quantity"`
	IsAlreadyRequestedOrBorrowed bool   `json:"isAlreadyRequestedOrBorrowed"`
	Status                       string `json:"status"`
}

type Request struct {
	RequestID  int    `json:"requestID"`
	UserID     int    `json:"userid"`
	BookID     int    `json:"bookid"`
	Username   string `json:"username"`
	Bookname   string `json:"bookname"`
	Author     string `json:"author"`
	IsAccepted bool   `json:"isAccepted"`
	IsBorrowed bool   `json:"isBorrowed"`
}

type BookList struct {
	Books []Book `json:"books"`
	Name  string `json:"name"`
}

type UserList struct {
	Users []User `json:"users"`
	Name  string `json:"name"`
}

type RequestList struct {
	Requests []Request `json:"requests"`
	Name     string    `json:"name"`
}
