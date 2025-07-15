package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func parseInput(input string) (time.Time, error) {
	if strings.ToLower(input) == "now" {
		return time.Now(), nil
	}

	// Try to parse as Unix timestamp
	if timestamp, err := strconv.ParseInt(input, 10, 64); err == nil {
		// Check if it's likely a Unix timestamp (reasonable range)
		if timestamp > 946684800 && timestamp < 4102444800 { // Between 2000 and 2100
			return time.Unix(timestamp, 0), nil
		}
	}

	// Try to parse as full timestamp (YYYYMMDDHHmmss)
	if len(input) == 14 {
		return time.Parse("20060102150405", input)
	}

	// Try to parse as date only (YYYYMMDD)
	if len(input) == 8 {
		// Append 010101 for 01:01:01
		fullTimestamp := input + "010101"
		return time.Parse("20060102150405", fullTimestamp)
	}

	// Try to parse as YYYY-MM-DD format
	if len(input) == 10 && strings.Contains(input, "-") {
		t, err := time.Parse("2006-01-02", input)
		if err == nil {
			// Set time to 01:01:01
			return time.Date(t.Year(), t.Month(), t.Day(), 1, 1, 1, 0, t.Location()), nil
		}
	}

	// Try to parse as ISO 8601 format
	if t, err := time.Parse(time.RFC3339, input); err == nil {
		return t, nil
	}

	// Try common datetime formats
	formats := []string{
		"2006-01-02 15:04:05",
		"2006/01/02 15:04:05",
		"02-01-2006 15:04:05",
		"02/01/2006 15:04:05",
	}

	for _, format := range formats {
		if t, err := time.Parse(format, input); err == nil {
			return t, nil
		}
	}

	return time.Time{}, fmt.Errorf("invalid date format: %s", input)
}

func generateFormats(t time.Time) {
	// YYYYMMDDHHmmss format
	timestamp := t.Format("20060102150405")
	wf.NewItem(timestamp).
		Subtitle("YYYYMMDDHHmmss format").
		Arg(timestamp).
		Valid(true)

	// Unix timestamp
	unixTimestamp := t.Unix()
	wf.NewItem(fmt.Sprintf("%d", unixTimestamp)).
		Subtitle("Unix timestamp (seconds)").
		Arg(fmt.Sprintf("%d", unixTimestamp)).
		Valid(true)

	// ISO 8601
	iso8601 := t.Format(time.RFC3339)
	wf.NewItem(iso8601).
		Subtitle("ISO 8601 format").
		Arg(iso8601).
		Valid(true)

	// Human readable
	humanReadable := t.Format("2006-01-02 15:04:05")
	wf.NewItem(humanReadable).
		Subtitle("YYYY-MM-DD HH:mm:ss").
		Arg(humanReadable).
		Valid(true)

	// Date only
	dateOnly := t.Format("2006-01-02")
	wf.NewItem(dateOnly).
		Subtitle("Date only (YYYY-MM-DD)").
		Arg(dateOnly).
		Valid(true)

	// RFC822
	rfc822 := t.Format(time.RFC822)
	wf.NewItem(rfc822).
		Subtitle("RFC822 format").
		Arg(rfc822).
		Valid(true)
}

func run() {
	args := wf.Args()
	var input string
	
	if len(args) == 0 {
		input = "now"
	} else {
		input = args[0]
	}
	
	targetTime, err := parseInput(input)
	if err != nil {
		wf.NewItem("Error").
			Subtitle(err.Error()).
			Valid(false)
		wf.SendFeedback()
		return
	}

	generateFormats(targetTime)
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}