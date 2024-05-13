package routes

import (
	"net/http"
	"backend/controller"
)





func AuthRoutes(mux *http.ServeMux){
//Login route
mux.HandleFunc("/login", controller.Login)
}