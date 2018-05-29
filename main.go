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

func volume() error {
	buffer := bytes.NewBufferString("")
	command := exec.Command("pamixer", "--get-mute")
	command.Stdout = buffer
	command.Run()
	if strings.Contains(buffer.String(), "true") {
		fmt.Print("🔇")
		return nil
	}
	buffer = bytes.NewBufferString("")
	command = exec.Command("pamixer", "--get-volume")
	command.Stdout = buffer
	command.Run()
	fmt.Print("🔊 " + strings.Replace(buffer.String(), "\n", "", -1) + "%")
	return nil
}

func battery() error {
	symbols := make(map[string]string)
	output := make([]string, 0)
	buffer := bytes.NewBufferString("")
	command := exec.Command("acpi")
	rePercent := regexp.MustCompile("[0-9]{1,}%")
	reStatus := regexp.MustCompile("Unknown|Charging|Full|Discharging")

	command.Stdout = buffer
	symbols["Full"] = "⚡"
	symbols["Charging"] = "🔌"
	symbols["Discharging"] = "🔋"
	symbols["Unknown"] = "❓"

	if err := command.Run(); err != nil {
		return err
	}

	for index, line := range strings.Split(buffer.String(), "\n") {
		status := reStatus.Find([]byte(line))
		percentage := rePercent.Find([]byte(line))
		statusSymbol, _ := symbols[string(status)]

		if len(percentage) == 0 {
			continue
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
	case "volume":
		if err := volume(); err != nil {
			fmt.Fprintf(os.Stderr, "ERROR: %s", err.Error())
		}
	default:
		// TODO: Print possible args
	}
}
