package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/gorilla/mux"
)

func main() {

	r := chi.NewRouter()

	// Configurando as opções do CORS
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Permitir solicitações de qualquer origem
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300, // Tempo em segundos para a informação do preflight ser armazenada em cache
	})
	// Aplicar as configurações de CORS para todas as rotas
	r.Use(cors.Handler)
	r.Get("/teste", handleTeste)
	r.Get("/teste/{valor}", handleTesteParam)
	r.Get("/exercicio3", handleQuery)
	r.Post("/formulario", handleFormulario)

	// Iniciando o servidor na porta 8081
	fmt.Println("Servidor escutando na porta :8081")
	err := http.ListenAndServe(":8081", r)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor:", err)
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

type FormData struct {
	Texto       string `json:"texto"`
	Inteiro     int    `json:"inteiro"`
	Booleano    bool   `json:"booleano"`
	OpcaoSelect string `json:"opcaoSelect"`
	OpcaoRadio  string `json:"opcaoRadio"`
}

func handleFormulario(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var formData FormData
	if err := json.NewDecoder(r.Body).Decode(&formData); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(formData)
	fmt.Println(formData)
}
