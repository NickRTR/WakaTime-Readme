package github

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/NickRTR/WakaTime-Readme/cli"
	"github.com/NickRTR/WakaTime-Readme/stats"
)

func theme() (string, string) {
	theme := os.Getenv("THEME")

	var empty string
	var full string

	switch theme {
	case "block-green":
		empty = "⬜"
		full = "🟩"
	case "block-yellow":
		empty = "⬜"
		full = "🟨"
	case "block-red":
		empty = "⬜"
		full = "🟥"
	case "block-purple":
		empty = "⬜"
		full = "🟪"
	case "block-orange":
		empty = "⬜"
		full = "🟧"
	case "block-blue":
		empty = "⬜"
		full = "🟦"
	case "block-black":
		empty = "⬜"
		full = "⬛"

	case "circle-green":
		empty = "⚪"
		full = "🟢"
	case "circle-yellow":
		empty = "⚪"
		full = "🟡"
	case "circle-red":
		empty = "⚪"
		full = "🔴"
	case "circle-purple":
		empty = "⚪"
		full = "🟣"
	case "circle-orange":
		empty = "⚪"
		full = "🟠"
	case "circle-blue":
		empty = "⚪"
		full = "🔵"
	case "circle-black":
		empty = "⚪"
		full = "⚫"
	case "default":
		empty = "░"
		full = "█"
	default:
		empty = "░"
		full = "█"
	}

	return empty, full
}

func Format7DaysStats(langs stats.Languages) string {
	empty, full := theme()

	var graph string
	graph += "<h2>Last 7 Days</h2>"

	if len(langs) == 0 {
		graph += "No coding activity found for the last 7 days ⛱️."
	} else {
		for i, l := range langs {
			if i > 4 {
				break
			}
			percent := math.Round(l.Percent)
			graph += fmt.Sprintf("%-15s %15s %s %6.2f %%</br>", l.Name, l.Text, strings.Repeat(full, int(percent/4))+strings.Repeat(empty, int(25-int(percent/4))), l.Percent)
		}
	}

	return graph
}

func FormatAllTimeStats(data stats.AllTimeStats) string {
	markdown := "<h2>All Time</h2>"
	markdown += fmt.Sprintf("<strong>%-20s</strong>%s", "Total Time Coded: ", data.Data.Text)

	timespan := data.Data.Range.End.Sub(data.Data.Range.Start).Hours()
	days := timespan / 24

	markdown += fmt.Sprintf("</br><strong>%-20s</strong>%d days", "Timespan: ", int(days))

	totalTime, err := strconv.ParseFloat(data.Data.Decimal, 64)
	if err != nil {
		cli.BrintErr(err.Error())
	} else {
		average := totalTime * 60 / days
		hours := int(average / 60)
		minutes := int(average) - (hours * 60)
		markdown += fmt.Sprintf("</br><strong>%-20s</strong>%d hr(s) %d min(s)", "Daily average: ", hours, minutes)
	}
	return markdown
}
