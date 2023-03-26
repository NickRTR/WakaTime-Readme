package github

import (
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/NickRTR/WakaTime-Readme/stats"
)

func Create7DaysGraph(langs stats.Languages) string {
	theme := os.Getenv("THEME")

	var empty string
	var done string

	switch theme {
	case "block-green":
		empty = "⬜"
		done = "🟩"
	case "block-yellow":
		empty = "⬜"
		done = "🟨"
	case "block-red":
		empty = "⬜"
		done = "🟥"
	case "block-purple":
		empty = "⬜"
		done = "🟪"
	case "block-orange":
		empty = "⬜"
		done = "🟧"
	case "block-blue":
		empty = "⬜"
		done = "🟦"
	case "block-black":
		empty = "⬜"
		done = "⬛"

	case "circle-green":
		empty = "⚪"
		done = "🟢"
	case "circle-yellow":
		empty = "⚪"
		done = "🟡"
	case "circle-red":
		empty = "⚪"
		done = "🔴"
	case "circle-purple":
		empty = "⚪"
		done = "🟣"
	case "circle-orange":
		empty = "⚪"
		done = "🟠"
	case "circle-blue":
		empty = "⚪"
		done = "🔵"
	case "circle-black":
		empty = "⚪"
		done = "⚫"
	case "default":
		empty = "░"
		done = "█"
	default:
		empty = "░"
		done = "█"
	}

	var graph string
	graph += "<h2>Last 7 Days</h2>"

	for i, l := range langs {
		if i > 4 {
			break
		}
		percent := math.Round(l.Percent)
		graph += fmt.Sprintf("%-15s %15s %s %5.2f %%</br>", l.Name, l.Text, strings.Repeat(done, int(percent/4))+strings.Repeat(empty, int(25-int(percent/4))), l.Percent)
	}

	return graph
}

func CreateAllTimeData(data stats.AllTimeStats) string {
	markdown := "<h2>All Time</h2>"
	markdown += fmt.Sprintf("<strong>Total Time Coded: </strong>%s", data.Data.Text)
	return markdown
}
