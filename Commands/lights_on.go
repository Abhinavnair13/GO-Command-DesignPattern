package commands

import receiver "Light-Bulb/Receiver"

type LightOn struct {
	Light *receiver.Light
}

func (l *LightOn) Execute() {
	l.Light.On()
}
func (l *LightOn) UnExecute() {
	l.Light.Off()
}
