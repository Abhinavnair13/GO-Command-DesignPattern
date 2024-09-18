package main

import (
	commands "Light-Bulb/Commands"
	invoker "Light-Bulb/Invoker"
	receiver "Light-Bulb/Receiver"
	"fmt"
)

func setupCommands(light *receiver.Light) (*commands.LightOn, *commands.LightOff) {
	lightOn := &commands.LightOn{Light: light}
	lightOff := &commands.LightOff{Light: light}
	return lightOn, lightOff
}

func performCommands(remote *invoker.RemoteControl, lightOn *commands.LightOn, lightOff *commands.LightOff) {
	fmt.Println("Buttons for lights ON switch")
	remote.SetCommand(lightOn)
	remote.PressButton()
	remote.PressUndoButton()
	fmt.Println("*****************************")
	fmt.Println("Buttons for lights OFF switch")
	remote.SetCommand(lightOff)
	remote.PressButton()
	remote.PressUndoButton()
}

func main() {
	// Receiver
	light := &receiver.Light{}

	// Commands
	lightOn, lightOff := setupCommands(light)

	// Invoker
	remote := &invoker.RemoteControl{}

	// Perform commands
	performCommands(remote, lightOn, lightOff)
}
