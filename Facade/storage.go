package Facade

import "fmt"

type Storage struct {
	engine string
}

func NewStorage(e string) Storage {
	return Storage{engine:e}
}

func (s *Storage) Save(comment string){
	fmt.Printf("register%s\n",comment)
}
