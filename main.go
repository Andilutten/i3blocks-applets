package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"
)

func clock() {
	t := time.Now()
	hour := string(t.Hour())
	minute := string(t.Minute())
	if len(hour) < 2 {
		hour = "0" + hour
	}
	if len(minute) < 2 {
		minute = "0" + minute
	}
	fmt.Fprintf(os.Stdout, "🕓 %s:%s", hour, minute)
}

func date() {
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

func help() {
	fmt.Println("Use with one of the following commands")
	fmt.Println("\ttime - Prints a clock")
	fmt.Println("\tdate - Shows the current date")
	fmt.Println("\tbattery - Shows the status of every battery on the device")
	fmt.Println("\tvolume - Prints the current volume or mute")
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	arg := os.Args[1]
	switch arg {
	case "time":
		clock()
	case "date":
		date()
	case "battery":
		battery()
	case "volume":
		volume()
	default:
		help()
	}
}
