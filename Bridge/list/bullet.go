package list

import "fmt"

type Bullet struct {
	Bullet rune
}

func NewBullet(b rune) *Bullet {
	return &Bullet{Bullet: b}
}

func (b *Bullet) Print(todos []string) {
	for _, v := range todos {
		fmt.Println("\t", string(b.Bullet), v)
	}
}
