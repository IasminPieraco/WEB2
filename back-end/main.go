package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/teste", handleCORS(handleTeste))
	r.HandleFunc("/teste/{valor}", handleCORS(handleTesteParam))
	r.HandleFunc("/exercicio3", handleCORS(handleQuery))
	// Iniciando o servidor na porta 8081
	fmt.Println("Servidor escutando na porta :8081")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
	}
}

func handleCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	}
}

// handleTeste é o handler para a nova rota /teste
func handleTeste(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Rota /teste executou com sucesso!")
}

// handleTesteParam é o handler para a rota /teste/{param}
func handleTesteParam(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	param := vars["valor"]
	fmt.Fprintf(w, "Rota /teste/%s executou com sucesso!", param)

}

// handleQuery é o handler para a rota /testeQuery
func handleQuery(w http.ResponseWriter, r *http.Request) {
	// Obter os valores da query
	valores := r.URL.Query()
	valor := valores.Get("valor")
	quantidade := valores.Get("quantidade")
	resposta := fmt.Sprintf("Rota /exercicio3 executou com sucesso recebendo o valor %s e quantidade %s!", valor, quantidade)
	fmt.Fprintf(w, resposta)
}
