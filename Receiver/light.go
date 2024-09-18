package receiver

import "fmt"

type Light struct {
	Brightness int
}

func (l *Light) On() {
	fmt.Println("Light is On")
}
func (l *Light) Off() {
	fmt.Println("Light is Off")
}
func (l *Light) DimLights(amt int) {
	if l.Brightness > 1 {
		l.Brightness -= amt
		fmt.Printf("Brightness set to %d %%\n", l.Brightness)
	} else {
		fmt.Printf("Brightness level is minimum : %d %%\n", l.Brightness)
	}
}
func (l *Light) BrightenLights(amt int) {
	if l.Brightness > 1 {
		l.Brightness += amt
		fmt.Printf("Brightness set to %d %%\n", l.Brightness)
	} else {
		fmt.Printf("Brightness level is minimum : %d %%\n", l.Brightness)
	}
}
