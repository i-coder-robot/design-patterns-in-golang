package main

import (
	"fmt"
	"github.com/i-coder-robot/design-patterns-in-golang/Decorator"
	"log"
	"os"
)

func main(){
	fmt.Println(Decorator.Pi(1000))
	fmt.Println(Decorator.Pi(50000))

	f:=Decorator.Wraplogger(Decorator.Pi,log.New(os.Stdout,"test",1))

	f(100000)
}
