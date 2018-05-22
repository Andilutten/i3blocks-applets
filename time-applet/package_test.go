package timeapplet

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

var ttime = time.Now()

func GetTimeDependency() time.Time {
	return ttime
}

func TestApplet(t *testing.T) {
	t.Run("HappyPath", func(t *testing.T) {
		expected := fmt.Sprintf("🕓 %v:%v", ttime.Hour(), ttime.Minute())
		out := bytes.NewBufferString("")

		TimeApp{
			GetTime: GetTimeDependency,
			Out:     out,
		}.Run()

		if out.String() != expected {
			t.Errorf("Got output: %s, expected: %s", out.String(), expected)
		}
	})
}
