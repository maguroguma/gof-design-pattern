package command

type Position struct {
	x, y int
}

func NewPosition(x, y int) *Position {
	p := new(Position)
	p.x, p.y = x, y
	return p
}
