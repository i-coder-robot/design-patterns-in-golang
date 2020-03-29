package Proxy

import (
	"fmt"
	"testing"
	"time"
)

var loc Proxy

func TestNew(t *testing.T) {
	loc = NewLocal()

	GetByID(2)
	GetByID(2)
	GetByID(1)
	GetByID(2)
	GetByID(3)
	GetByID(2)
	GetByID(1)
	GetAll()
}

func GetByID(id int32){
	start := time.Now()
	fmt.Printf("%+v", loc.GetByID(id))
	elapsed := time.Since(start)
	fmt.Printf("Tiempo usado: %v\n", elapsed)
}

func GetAll(){
	start := time.Now()
	fmt.Printf("%+v", loc.GetAll())
	elapsed := time.Since(start)
	fmt.Printf("Tiempo usado: %v\n", elapsed)
}
