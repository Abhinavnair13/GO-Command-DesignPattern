package invoker

import commands "Light-Bulb/Commands"

type RemoteControl struct {
	command commands.Command
}

// remote control can be of any lights, this is used to say that hey i am
// using this remote for this particular light. It is like a setter
func (r *RemoteControl) SetCommand(command commands.Command) {
	r.command = command

}
func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

func (r *RemoteControl) PressUndoButton() {
	r.command.UnExecute()
}
