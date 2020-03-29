package Proxy

type Local struct {
	Remote *Data
	cache Books
}

func NewLocal() Proxy{
	return &Local{
		Remote: New("https://alguno.com", 8080, "usuario", "123456"),
		cache:  make(Books, 0),
	}
}

func (l *Local) GetByID(id int32) Book{
	for _,v :=range l.cache{
		if v.ID==id{
			return v
		}
	}
	b:=l.Remote.ByID(id)
	l.cache=append(l.cache,b)
	return b
}

func (l *Local) GetAll() Books {
	return l.Remote.All()
}