package Bridage

import "testing"

func TestCircle_Draw(t *testing.T) {
	redCircle := Circle{}
	redCircle.Constructor(100, 100, 10, &RedCircle{})
	yellowCircle := Circle{}
	yellowCircle.Constructor(200, 200, 20, &YellowCircle{})
	redCircle.Cook()
	yellowCircle.Cook()
}
