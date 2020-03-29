package Proxy

type Proxy interface {
	GetByID(id int32) Book
	GetAll() Books
}