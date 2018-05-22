package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	timeapplet "github.com/Andilutten/i3blocks-applets/time-applet"
)

func appDate() {
	t := time.Now()
	s := fmt.Sprintf("📆 %v-%v-%v", t.Day(), t.Month().String(), t.Year())
	fmt.Fprintf(os.Stdout, s)
}

func appBattery() {
	out := bytes.NewBufferString("")
	cmd := exec.Command("acpi")
	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	var status string
	items := strings.Split(out.String(), " ")
	switch items[2] {
	case "Full,":
		status = "⚡"
	case "Charging,":
		status = "🔌"
	case "Discharging":
		status = "🔋"
	default:
		status = "❓"
	}

	fmt.Fprintf(os.Stdout, "%s %s", status, items[3])
}

func appWeather() {
	out := bytes.NewBufferString("")
	cmd := exec.Command("curl", "wttr.in/~Sweden+Trollhättan")

	cmd.Stdout = out
	err := cmd.Run()
	if err != nil {
		panic(err)
	}

	line := strings.Split(out.String(), "\n")[12]
	re := regexp.MustCompile("[0-9][0-9]")

	hits := re.FindAll([]byte(line), -1)
	for _, hit := range hits {
		fmt.Println(string(hit))
	}
}

func main() {
	if len(os.Args) < 2 {
		// TODO: Print possible args
		return
	}

	arg := os.Args[1]
	switch arg {
	case "time":
		timeapplet.TimeApp{
			GetTime: time.Now,
			Out:     os.Stdout,
		}.Run()
	case "weather":
		appWeather()
	case "date":
		appDate()
	case "battery":
		appBattery()
	default:
		// TODO: Print possible args
	}
}
