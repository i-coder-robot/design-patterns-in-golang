package Command

import "testing"

func TestPerson_Talk(t *testing.T) {
	laowang := NewPerson("wang", NewCommand(nil, nil))
	laozhang := NewPerson("zhang", NewCommand(&laowang, laowang.Listen))
	laofeng := NewPerson("feng", NewCommand(&laozhang, laozhang.Gossip))
	laoding := NewPerson("ding", NewCommand(&laofeng, laofeng.PassOn))

	laoding.Talk()
}
