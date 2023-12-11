package webserver

import "testing"

func TestConsultaCepNoResult(t *testing.T) {
	cep := "99001555"
	endereco := BuscarCep(cep)

	if endereco.CEP != "" {
		t.Fatalf("CEP não encontrado %s", endereco.CEP)
	}
}

func TestConsultaCepWithResult(t *testing.T) {
	cep := "75258-861"
	endereco := BuscarCep(cep)

	if endereco.CEP != cep {
		t.Fatalf("CEP não encontrado %s", endereco.CEP)
	}
	t.Logf("CEP encontrado %s", endereco.CEP)
}
