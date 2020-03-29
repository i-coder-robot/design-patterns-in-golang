package Prototype

import (
	"fmt"
	"testing"
)

func TestPrototype_Clone(t *testing.T) {
	color:="red"
	phones:=map[string]string{"home":"123456","work":"7890123"}
	p1:=prototype{"欢喜哥", 24, []string{"Mike", "Tom"}, &color, phones}
	p2:=p1.Clone()

	p2.age=20
	p2.name="小明"
	p2.friends[0] = "小红"
	p2.friends[1] = "老王"
	color = "azul"
	phones["home"] = "147258"

	fmt.Println(p1.String())
	fmt.Println(p2.String())
}
