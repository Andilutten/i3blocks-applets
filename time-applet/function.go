package timeapplet

import "fmt"

// Run the applet
func (app TimeApp) Run() {
	t := app.GetTime()

	s := fmt.Sprintf("🕓 %v:%v", t.Hour(), t.Minute())
	fmt.Fprint(app.Out, s)
}
