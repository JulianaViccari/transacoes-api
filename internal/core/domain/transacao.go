package domain

import (
	"errors"
	"time"
)

type Transacao struct {
	Valor float64
	DataHora time.Time
}

func NewTransacao(valor float64, dataHora time.Time) Transacao{
	return Transacao{
		Valor: valor,
		DataHora: dataHora,
	}
}

func (t Transacao) isValidDate() bool{
	return t.DataHora.Compare(time.Now()) == -1 
}

func (t Transacao) isValidValue() bool{
	return t.Valor >= 0
}

func (t Transacao) Validate() error{
	if !t.isValidDate(){
		return errors.New("data da transação inválida")
	}

	if !t.isValidValue(){
		return errors.New("valor da transação inválido")
	}

	return nil
}

