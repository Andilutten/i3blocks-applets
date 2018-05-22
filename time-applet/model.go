package timeapplet

import (
	"io"
	"time"
)

// A TimeFunc is a function
// that returns a Time object
type TimeFunc func() time.Time

// A TimeApp is an object
// containing all data for
// the time applet
type TimeApp struct {
	GetTime TimeFunc
	Out     io.Writer
}
