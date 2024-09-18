package receiver

import "fmt"

type Light struct{}

func (l *Light) On() {
	fmt.Println("Light is On")
}
func (l *Light) Off() {
	fmt.Println("Light is Off")
}
