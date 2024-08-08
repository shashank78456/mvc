package api

import (
	"github.com/gorilla/mux"
	"github.com/shashank78456/mvc/pkg/controller"
	"net/http"
)

func Start() {
	router := mux.NewRouter()

	publicRouter := router.PathPrefix("/").Subrouter()
	staticFiles := http.FileServer(http.Dir("./public/"))
	publicRouter.PathPrefix("/public/").Handler(http.StripPrefix("/public/", staticFiles))

	authRouter := router.PathPrefix("/").Subrouter()
	authRouter.Use(controller.Authenticator)
	authRouter.HandleFunc("/", controller.RenderLogin).Methods("GET")
	authRouter.HandleFunc("/", controller.HandleLogin).Methods("POST")
	authRouter.HandleFunc("/signup", controller.RenderSignup).Methods("GET")
	authRouter.HandleFunc("/signup", controller.HandleSignup).Methods("POST")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(controller.Authenticator)
	adminRouter.HandleFunc("/{page}", controller.RenderAdmin).Methods("GET")
	adminRouter.HandleFunc("/add_new_book", controller.AddNewBook).Methods("POST")
	adminRouter.HandleFunc("/add_book", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/remove_book", controller.RemoveBook).Methods("POST")
	adminRouter.HandleFunc("/delete_book", controller.DeleteBook).Methods("DELETE")
	adminRouter.HandleFunc("/handle_request", controller.HandleRequest).Methods("POST")

	superadminRouter := router.PathPrefix("/superadmin").Subrouter()
	superadminRouter.Use(controller.Authenticator)
	superadminRouter.HandleFunc("/{page}", controller.RenderSuperAdmin).Methods("GET")
	superadminRouter.HandleFunc("/add_new_book", controller.AddNewBook).Methods("POST")
	superadminRouter.HandleFunc("/add_book", controller.AddBook).Methods("POST")
	superadminRouter.HandleFunc("/remove_book", controller.RemoveBook).Methods("POST")
	superadminRouter.HandleFunc("/delete_book", controller.DeleteBook).Methods("DELETE")
	superadminRouter.HandleFunc("/handle_request", controller.HandleRequest).Methods("POST")
	superadminRouter.HandleFunc("/accept_admin", controller.AcceptAdmin).Methods("POST")
	superadminRouter.HandleFunc("/deny_admin", controller.DenyAdmin).Methods("POST")

	clientRouter := router.PathPrefix("/client").Subrouter()
	clientRouter.Use(controller.Authenticator)
	clientRouter.HandleFunc("/{page}", controller.RenderClient).Methods("GET")
	clientRouter.HandleFunc("/request_book", controller.RequestBook).Methods("POST")
	clientRouter.HandleFunc("/return_book", controller.ReturnBook).Methods("POST")
	clientRouter.HandleFunc("/admin_request", controller.AdminRequest).Methods("POST")

	http.ListenAndServe(":3000", router)
}
