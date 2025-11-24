package internal

import (
	"errors"
	"unicode"
)

type Cep struct {
	cep string
}

func NewCep(cep string) (*Cep, error) {
	if len(cep) != 8 {
		return nil, errors.New("invalid zipcode")
	}

	for _, r := range cep {
		if !unicode.IsDigit(r) {
			return nil, errors.New("invalid zipcode")
		}
	}

	return &Cep{
		cep: cep,
	}, nil
}

func (c *Cep) Get() string {
	return c.cep
}
