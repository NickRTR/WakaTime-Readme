package requests

import (
	"encoding/json"
	"fmt"

	"github.com/NickRTR/WakaTime-Readme/stats"
)

func Last7Days(token string) stats.Languages {
	res := stats.Request(fmt.Sprintf("https://wakatime.com/api/v1/users/current/stats/last_7_days?api_key=%s&scope=read_stats", token))

	var data stats.SevenDaysStats
	json.Unmarshal([]byte(res), &data)

	return data.Data.Languages
}
