package commands

type Command interface {
	Execute()
	UnExecute()
}
