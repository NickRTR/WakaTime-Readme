package requests

import (
	"encoding/json"
	"fmt"

	"github.com/NickRTR/WakaTime-Readme/stats"
)

func AllTime(token string) stats.AllTimeStats {
	res := stats.Request(fmt.Sprintf("https://wakatime.com/api/v1/users/current/all_time_since_today?api_key=%s", token))

	var data stats.AllTimeStats
	json.Unmarshal([]byte(res), &data)

	return data
}
