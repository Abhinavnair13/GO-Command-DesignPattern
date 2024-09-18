package commands

import receiver "Light-Bulb/Receiver"

type LightOff struct {
	Light *receiver.Light
}

func (l *LightOff) Execute() {
	l.Light.Off()
}
func (l *LightOff) UnExecute() {
	l.Light.On()
}
