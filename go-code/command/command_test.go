package command

import "testing"

func TestCommandPattern(t *testing.T) {
	macro := NewMacroCommand()

	macro.Append(NewDrawCommand(NewPosition(1, 1)))
	macro.Append(NewDrawCommand(NewPosition(2, 2)))

	println(macro.Execute())
	println(macro.Execute())

	macro.Undo()
	println(macro.Execute())
	println(macro.Execute())
}

func TestCommandPatternWithCompositePattern(t *testing.T) {
	macro := NewMacroCommand()
	macro.Append(NewDrawCommand(NewPosition(1, 1)))
	macro.Append(NewDrawCommand(NewPosition(2, 2)))

	com := NewMacroCommand()
	com.Append(NewDrawCommand(NewPosition(3, 3)))
	com.Append(NewDrawCommand(NewPosition(4, 4)))

	macro.Append(com)

	println(macro.Execute())

	macro.Undo()
	println(macro.Execute())
}
