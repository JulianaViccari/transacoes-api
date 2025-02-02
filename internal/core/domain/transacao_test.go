package domain

import(
	"testing"
	"time"
	"fmt"
)

func TestTansacaoValida(t *testing.T) {
	transacao := NewTransacao(50.00, time.Now())
	fmt.Println(transacao) 
	if err := transacao.Validate(); err != nil{
		t.Error("teste")
	}
}

func TestTransacaoInvalidaValorNegativo(t *testing.T) {
	transacao := NewTransacao(-50.00, time.Now())
	mensagemErro := "valor da transação inválido" 

	if err := transacao.Validate(); err == nil{
		t.Errorf("esperado mensagem de erro: %s", mensagemErro)
	}
	
	if err := transacao.Validate(); err != nil{
		if err.Error() != mensagemErro{
			t.Errorf("esperado mensagem de erro: %s", mensagemErro)
		}
	}
}

func TestTransacaoInvalidaDataFutura(t *testing.T) {
	transacao := NewTransacao(100.00, time.Now().Add(time.Minute*10))
	mensagemErro := "data da transação inválida"
	fmt.Println(transacao)

	if err := transacao.Validate(); err == nil {
		t.Errorf("esperado mensagem de erro: %s", mensagemErro)
	}
}

