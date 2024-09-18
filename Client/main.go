package main

import (
	commands "Light-Bulb/Commands"
	invoker "Light-Bulb/Invoker"
	receiver "Light-Bulb/Receiver"
	"fmt"
)

func setupCommands(light *receiver.Light) (*commands.LightOn, *commands.LightOff, *commands.LightBrightness) {
	lightOn := &commands.LightOn{Light: light}
	lightOff := &commands.LightOff{Light: light}
	lightBrightness := &commands.LightBrightness{Light: light, Brightness: light.Brightness}
	return lightOn, lightOff, lightBrightness
}

func performONOffCommands(remote *invoker.RemoteControl, lightOn *commands.LightOn, lightOff *commands.LightOff) {
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
func performBrightnessCommands(remote *invoker.RemoteControl, brightness *commands.LightBrightness) {
	fmt.Printf("Orginal brightness : %d %%\n", brightness.Brightness)
	remote.SetCommand(brightness)
	remote.PressButton()
	fmt.Println("Brightness lvl increased")
	remote.PressUndoButton()
	fmt.Println("Brightness lvl decreased")
}

func main() {
	// Receiver
	light := &receiver.Light{Brightness: 5}

	// Commands
	lightOn, lightOff, lightBrightness := setupCommands(light)

	// Invoker
	remote := &invoker.RemoteControl{}

	// Perform commands
	performONOffCommands(remote, lightOn, lightOff)
	fmt.Println("################################")
	performBrightnessCommands(remote, lightBrightness)

}
