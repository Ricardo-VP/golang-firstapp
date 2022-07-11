package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func index(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(w, "Método no permitido")
		return
	}

	fmt.Fprintf(w, "Hola mundo")
}

func suma(w http.ResponseWriter, r *http.Request) {
	numero1 := r.URL.Query()["numero1"]
	numero2 := r.URL.Query()["numero2"]

	numero1Float, _ := strconv.ParseFloat(numero1[0], 4)
	numero2Float, _ := strconv.ParseFloat(numero2[0], 4)
	resultado := numero1Float + numero2Float

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "OK",
		"resultado": fmt.Sprint(resultado),
	})
}

func resta(w http.ResponseWriter, r *http.Request) {
	numero1 := r.URL.Query()["numero1"]
	numero2 := r.URL.Query()["numero2"]

	numero1Float, _ := strconv.ParseFloat(numero1[0], 4)
	numero2Float, _ := strconv.ParseFloat(numero2[0], 4)
	resultado := numero1Float - numero2Float

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "OK",
		"resultado": fmt.Sprint(resultado),
	})
}

func multiplicacion(w http.ResponseWriter, r *http.Request) {
	numero1 := r.URL.Query()["numero1"]
	numero2 := r.URL.Query()["numero2"]

	numero1Float, _ := strconv.ParseFloat(numero1[0], 4)
	numero2Float, _ := strconv.ParseFloat(numero2[0], 4)
	resultado := numero1Float * numero2Float

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "OK",
		"resultado": fmt.Sprint(resultado),
	})
}

func division(w http.ResponseWriter, r *http.Request) {
	numero1 := r.URL.Query()["numero1"]
	numero2 := r.URL.Query()["numero2"]

	numero1Float, _ := strconv.ParseFloat(numero1[0], 4)
	numero2Float, _ := strconv.ParseFloat(numero2[0], 4)
	resultado := numero1Float / numero2Float

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":    "OK",
		"resultado": fmt.Sprint(resultado),
	})
}
