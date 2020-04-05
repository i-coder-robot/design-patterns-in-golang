package Command

import "testing"

func TestCommand_Execute(t *testing.T) {
	laowang := NewPerson("wang", NewCommand(nil, nil))
	laozhang := NewPerson("zhang", NewCommand(&laowang, laowang.Listen))
	loafeng := NewPerson("feng", NewCommand(&laozhang, laozhang.Buy))
	laoding := NewPerson("ding", NewCommand(&loafeng, loafeng.Cook))
	laoli := NewPerson("li", NewCommand(&laoding, laoding.Wash))

	laoli.Talk()

}
