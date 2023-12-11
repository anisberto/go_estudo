package webserver

import (
	"encoding/json"
	"net/http"

	"golang.org/x/exp/slog"
)

type Endereco struct {
	Bairro      string `json:"bairro"`
	CEP         string `json:"cep"`
	Complemento string `json:"complemento"`
	DDD         string `json:"ddd"`
	Localidade  string `json:"localidade"`
	Logradouro  string `json:"logradouro"`
	UF          string `json:"uf"`
}

type EnderecoPojo struct {
	Bairro      string `json:"Bairro"`
	CEP         string `json:"CEP"`
	Complemento string `json:"Complemento"`
	DDD         string `json:"Codigo DDD"`
	Localidade  string `json:"Cidade"`
	Logradouro  string `json:"Logradouro"`
	UF          string `json:"UF"`
}

var endereco Endereco

func BuscarCep(cep string) EnderecoPojo {

	url := structPathToSearchCep("json", cep)
	response, erro := http.Get(url)

	if erro != nil {
		slog.Info("Erro ao acessar a URL")
	} else {

		decoder := json.NewDecoder(response.Body)
		err := decoder.Decode(&endereco)

		if err != nil {
			slog.Info("Erro ao decodificar JSON:", err)
		} else {
			jsonString, _ := json.Marshal(endereco)
			slog.Info(string(jsonString))
			return constructEnderecoPojo(&endereco)
		}
	}
	return EnderecoPojo{}
}

func structPathToSearchCep(tipo string, cep string) string {
	baseUrl := "https://viacep.com.br/ws/"
	switch tipo {
	case "xml":
		return string(baseUrl + cep + "/xml/")
	case "text":
		return string(baseUrl + cep + "/text/")
	default:
		return string(baseUrl + cep + "/json/")
	}
}

func constructEnderecoPojo(enderecoEntity *Endereco) EnderecoPojo {
	if enderecoEntity == nil {
		return EnderecoPojo{}
	}
	return EnderecoPojo{
		Bairro:      enderecoEntity.Bairro,
		CEP:         enderecoEntity.CEP,
		Complemento: enderecoEntity.Complemento,
		DDD:         enderecoEntity.DDD,
		Localidade:  enderecoEntity.Localidade,
		Logradouro:  enderecoEntity.Logradouro,
		UF:          enderecoEntity.UF,
	}
}
