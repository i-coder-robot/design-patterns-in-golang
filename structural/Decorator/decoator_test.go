package Decorator

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPi(t *testing.T) {
	fmt.Println(Pi(5000))
	fmt.Println(Pi(10000))
	fmt.Println(Pi(50000))

}

func TestWrapLogger(t *testing.T) {
	f:= WrapLogger(Pi,log.New(os.Stdout,"test ",1))
	f(100000)
}