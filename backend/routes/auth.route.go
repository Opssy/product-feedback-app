package routes

import (
	"fmt"
	"net/http"
	"backend/controller"
)





func main(){
	mux := http.NewServeMux()

	//Login route
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w. WriteHeader(http.StatusMethodNotAllowed)
			fmt.Fprintf(w, "Method not allowed: %s", r.Method)
			return
		}
		controller.Login(w,r)
	})
}