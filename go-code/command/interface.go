package command

type command interface {
	Execute() string
}
