package Factory

import "fmt"

type Restaurant interface {
	GetFood()
}

type Donglaishun struct {
}

func (d *Donglaishun) GetFood() {
	fmt.Println("东来顺的饭菜准备继续")
}

type Qingfeng struct {
}

func (q *Qingfeng) GetFood() {
	fmt.Println("庆丰包子铺饭菜准备就绪")
}

func NewRestaurant(s string) Restaurant {
	switch s {
	case "d":
		return &Donglaishun{}
	case "q":
		return &Qingfeng{}
	}
	return nil
}
