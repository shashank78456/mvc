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

- MySQL

`mysql -u root -p < ./config/db.sql` and then enter password.

- Running Server

1. Execute `go run main.go`.

- Project Specs

The first user to signup has superadmin privileges.
Only superadmin can accept or deny clients seeking admin privileges.
Users cannot signup as admin.

Clients can borrow books by making a borrow request which needs to be approved by an admin.
Clients can return books and access their borrow history. Clients can request for admin privileges.

Admins can approve borrow requests, add new books in the catalog, make a book available for borrowing or disable borrowing of a book. 
