package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc(("/"), func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodGet {
			fmt.Fprintf(w, "Método no permitido")
			return
		}

		fmt.Fprintf(w, "Hola %s", "visitante")
	})

	srv := http.Server{
		Addr: ":8080",
	}

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
