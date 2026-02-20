// TESTE DE UNIDADE

package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
) 

type cenarioTeste struct {
	enderecoInserido string
	retornoEsperado string
}

func TestTipoDeEndereco(t *testing.T) {
	t.Parallel()

	cenarioTeste := []cenarioTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Paulista", "Avenida"},
		{"Rodovia dos Imigrantes", "Rodovia"},
		//{"Praça das Rodas", "Tipo Inválido"},
		{"Estrada Qualquer", "Estrada"},
		{"RUA DOS BOBOS", "Rua"},
		{"AVENIDA REBOUÇAS", "Avenida"},
		//{"", "Tipo Inválido"},
	}
	
	for _, cenario := range cenarioTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if retornoRecebido != cenario.retornoEsperado{
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
			retornoRecebido,
			cenario.retornoEsperado,
		)
		}
	}

}

func TestQualquer(t *testing.T) {
	t.Parallel()
	if 1 > 2 {
		t.Errorf("Teste quebrou!")
	}
}