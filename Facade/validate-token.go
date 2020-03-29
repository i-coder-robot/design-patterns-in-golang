package Facade

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("no login")

type ValidatorToken struct {
	token string
}

func NewValidatorToken(t string) ValidatorToken{
	return ValidatorToken{token:t}
}

func (vt *ValidatorToken) Validate() error{
	if vt.token!="token-validate"{
		return ErrTokenNotValid
	}

	fmt.Println("token validate")
	return nil
}