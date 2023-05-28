package routes

import (
	"fmt"
	"net/http"
)

func IndexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to my REST API")
}