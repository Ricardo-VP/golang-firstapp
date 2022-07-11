package server

import (
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/suma", suma)
	http.HandleFunc("/resta", resta)
	http.HandleFunc("/multiplicacion", multiplicacion)
	http.HandleFunc("/division", division)
}
