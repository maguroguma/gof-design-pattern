package command

type MacroCommand struct {
	commands []command
}

func NewMacroCommand() *MacroCommand {
	mc := new(MacroCommand)
	mc.commands = []command{}
	return mc
}

func (mc *MacroCommand) Execute() string {
	res := ""
	for _, command := range mc.commands {
		res += command.Execute() + "\n"
	}
	return res
}

func (mc *MacroCommand) Append(command command) {
	mc.commands = append(mc.commands, command)
}

func (mc *MacroCommand) Undo() {
	if len(mc.commands) > 0 {
		mc.commands = mc.commands[:len(mc.commands)-1]
	}
}

func (mc *MacroCommand) Clear() {
	mc.commands = []command{}
}
