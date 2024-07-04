## Library Management System

- Setting up

1. Clone the repo using `git clone https://github.com/shashank78456/mvc.git`.
2. Go to root directory of the repo. `cd mvc`.
3. Execute the following commands:
```
cp .env.sample .env
go mod vendor
go mod tidy
```

- Database

1. Execute `mysql -u root -p -e "DROP DATABASE IF EXISTS Library; CREATE DATABASE Library;"`

2. From project root directory execute `migrate -path ./database/migrations -database "mysql://username:password@(127.0.0.1:3306)/Library" up`

Replace username and password with your own.

- Running Server

1. Execute `go run main.go`.

- Unit Tests

There are two test files:
1. `users_test.go` which tests `func GetUserType(string) (int, error)` of users model
2. `auth_test.go` contains benchmark tests for `func createToken(string, string) (string, error)` , `func HashPassword(string, string) (string, error)` , `func IsPasswordValid(string, string) (bool)` of auth.go in controller

To run `users_test.go`:
```
cd pkg/models
go test -v
```

To run `auth_test.go`: 
```
cd pkg/controller
go test -bench=.
```
- Project Specs

The first user to signup has superadmin privileges.
Only superadmin can accept or deny clients seeking admin privileges.
Users cannot signup as admin.

Clients can borrow books by making a borrow request which needs to be approved by an admin.
Clients can return books and access their borrow history. Clients can request for admin privileges.

Admins can approve borrow requests, add new books in the catalog, make a book available for borrowing or disable borrowing of a book. 
