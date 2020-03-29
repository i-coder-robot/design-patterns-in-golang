package Facade

import (
	"fmt"
	"log"
	"testing"
)

func TestStorage_Save(t *testing.T) {
	showInit()
	token:="token-validate"
	user := "user-blog"
	to := "abc@gmail.com"
	comment := "video :)"

	f:= NewFacade(to,comment,token,user)
	err := f.Comment()
	if err != nil {
		log.Fatal(err)
	}
	f.Notify()
	showFinish()
}

func showInit() {
	fmt.Println(`
	**************************
	* Bienvenido al programa *
	**************************
	`)
}

func showFinish() {
	fmt.Println(`
	**************************
	*  Gracias por utilizar  *
	**************************
	`)
}
