package commands

import receiver "Light-Bulb/Receiver"

type LightBrightness struct {
	Light      *receiver.Light
	Brightness int
}

func (l *LightBrightness) Execute() {
	l.Light.BrightenLights(l.Brightness)
}
func (l *LightBrightness) UnExecute() {
	l.Light.DimLights(l.Brightness)
}
