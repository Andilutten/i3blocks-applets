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

func battery() error {
	output := make([]string, 0)
	buffer := bytes.NewBufferString("")
	command := exec.Command("acpi")
	rePercent := regexp.MustCompile("[0-9]{1,}%")
	reStatus := regexp.MustCompile("Charging|Full|Discharging")

	command.Stdout = buffer

	if err := command.Run(); err != nil {
		return err
	}

	for index, line := range strings.Split(buffer.String(), "\n") {
		status := reStatus.Find([]byte(line))
		percentage := rePercent.Find([]byte(line))
		var statusSymbol string

		if len(percentage) == 0 {
			continue
		}

		switch string(status) {
		case "Full":
			statusSymbol = "⚡"
		case "Charging":
			statusSymbol = "🔌"
		case "Discharging":
			statusSymbol = "🔋"
		default:
			statusSymbol = "❓"
		}

		output = append(output, fmt.Sprintf("Battery %v: %s %s", index+1, percentage, statusSymbol))
	}
	fmt.Fprint(os.Stdout, strings.Join(output, ", "))
	return nil
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
		battery()
	default:
		// TODO: Print possible args
	}
}
