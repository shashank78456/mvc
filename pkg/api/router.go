package api

import(
	"net/http"
	"github.com/shashank78456/mvc/pkg/controller"
	"github.com/gorilla/mux"
)

func Start() {
	router := mux.NewRouter()

	router.HandleFunc("/", controller.RenderLogin).Methods("GET")
	router.HandleFunc("/", controller.HandleLogin).Methods("POST")

	router.HandleFunc("/signup", controller.RenderSignup).Methods("GET")
	router.HandleFunc("/signup", controller.HandleSignup).Methods("POST")

	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(controller.Authenticator)
	adminRouter.HandleFunc("/{page}", controller.RenderAdmin).Methods("GET")
	adminRouter.HandleFunc("/add_new_book", controller.AddNewBook).Methods("POST")
	adminRouter.HandleFunc("/add_book", controller.AddBook).Methods("POST")
	adminRouter.HandleFunc("/remove_book", controller.RemoveBook).Methods("POST")
	adminRouter.HandleFunc("/delete_book", controller.DeleteBook).Methods("DELETE")
	adminRouter.HandleFunc("/accept_request", controller.AcceptRequest).Methods("POST")
	adminRouter.HandleFunc("/accept_admin", controller.AcceptAdmin).Methods("POST")
	adminRouter.HandleFunc("/deny_admin", controller.DenyAdmin).Methods("POST")

	clientRouter := router.PathPrefix("/client").Subrouter()
	clientRouter.Use(controller.Authenticator)
	clientRouter.HandleFunc("/{page}", controller.RenderClient).Methods("GET")
	clientRouter.HandleFunc("/request_book", controller.RequestBook).Methods("POST")
	clientRouter.HandleFunc("/return_book", controller.ReturnBook).Methods("POST")
	clientRouter.HandleFunc("/admin_request", controller.AdminRequest).Methods("POST")

	http.ListenAndServe(":3000", router)
}