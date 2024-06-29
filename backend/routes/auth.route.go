package routes

import (
	"net/http"
	"backend/controller"
)

func AuthRoutes(mux *http.ServeMux){
//Login route
mux.HandleFunc("/login", controller.Login)
mux.HandleFunc("/signup", controller.Signup)
mux.HandleFunc("/home", controller.Home)
// mux.HandleFunc("/edit", controller.Edit)
}
