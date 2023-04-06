package server

import (
	"encoding/json"
	"fmt"
	"funds/api"
	"net/http"
	"strings"
)

type ResultadoBusca struct {
	Busca  string `json:"busca"`
	Status string `json:"status"`
}

// FundsServer inicia o servidor na porta especificada.
func FundsServer() {
	http.HandleFunc("/v1", func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Query().Get("fundo") == "" {
			resultadoBusca := "{}"

			json, _ := json.MarshalIndent(resultadoBusca, "", "    ")

			w.Write(json)
			return
		}

		fundo := r.URL.Query().Get("fundo")

		structedFund, err := api.FundsApi(fundo)
		if err != nil {

			resultadoBusca := ResultadoBusca{
				Busca:  strings.ToUpper(fundo),
				Status: "Fundo não encontrado",
			}

			json, _ := json.MarshalIndent(resultadoBusca, "", "    ")

			w.Write(json)
			println("Erro ao obter fundo: ", err)
			return
		}

		json, err := json.MarshalIndent(structedFund, "", "	")
		if err != nil {
			return
		}

		w.Header().Set("Content-Type", "application/json")

		w.Write(json)
	})

	fmt.Println("Rodando na porta 8080...")
	http.ListenAndServe(":8080", nil)
}
