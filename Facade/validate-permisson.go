package Facade

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("no authorized")

type ValidatorPermission struct {
	userID string
}

func NewValidatorPermission(id string) ValidatorPermission {
	return ValidatorPermission{userID:id}
}

func (vp *ValidatorPermission) Validate() error {
	if vp.userID!="user-blog"{
		return ErrPermissionNotValid
	}
	fmt.Println("user authorized")
	return nil
}