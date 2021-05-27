package command

import "strconv"

type DrawCommand struct {
	Position *Position
}

func NewDrawCommand(p *Position) *DrawCommand {
	dc := new(DrawCommand)
	dc.Position = p
	return dc
}

func (dc *DrawCommand) Execute() string {
	return strconv.Itoa(dc.Position.x) + "." + strconv.Itoa(dc.Position.y)
}
