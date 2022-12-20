// API Response structs

package main

import "time"

type stats struct {
	Data struct {
		BestDay struct {
			CreatedAt    time.Time `json:"created_at"`
			Date         string    `json:"date"`
			ID           string    `json:"id"`
			ModifiedAt   time.Time `json:"modified_at"`
			Text         string    `json:"text"`
			TotalSeconds float64   `json:"total_seconds"`
		} `json:"best_day"`
		Categories []struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"categories"`
		CreatedAt                          time.Time `json:"created_at"`
		DailyAverage                       float64   `json:"daily_average"`
		DailyAverageIncludingOtherLanguage float64   `json:"daily_average_including_other_language"`
		DaysIncludingHolidays              int       `json:"days_including_holidays"`
		DaysMinusHolidays                  int       `json:"days_minus_holidays"`
		Dependencies                       []struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"dependencies"`
		Editors []struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"editors"`
		End                                             time.Time `json:"end"`
		Holidays                                        int       `json:"holidays"`
		HumanReadableDailyAverage                       string    `json:"human_readable_daily_average"`
		HumanReadableDailyAverageIncludingOtherLanguage string    `json:"human_readable_daily_average_including_other_language"`
		HumanReadableRange                              string    `json:"human_readable_range"`
		HumanReadableTotal                              string    `json:"human_readable_total"`
		HumanReadableTotalIncludingOtherLanguage        string    `json:"human_readable_total_including_other_language"`
		ID                                              string    `json:"id"`
		IsAlreadyUpdating                               bool      `json:"is_already_updating"`
		IsCodingActivityVisible                         bool      `json:"is_coding_activity_visible"`
		IsIncludingToday                                bool      `json:"is_including_today"`
		IsOtherUsageVisible                             bool      `json:"is_other_usage_visible"`
		IsStuck                                         bool      `json:"is_stuck"`
		IsUpToDate                                      bool      `json:"is_up_to_date"`
		IsUpToDatePendingFuture                         bool      `json:"is_up_to_date_pending_future"`
		Languages                                       []struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"languages"`
		Machines []struct {
			Decimal       string  `json:"decimal"`
			Digital       string  `json:"digital"`
			Hours         int     `json:"hours"`
			MachineNameID string  `json:"machine_name_id"`
			Minutes       int     `json:"minutes"`
			Name          string  `json:"name"`
			Percent       float64 `json:"percent"`
			Text          string  `json:"text"`
			TotalSeconds  float64 `json:"total_seconds"`
		} `json:"machines"`
		ModifiedAt       time.Time `json:"modified_at"`
		OperatingSystems []struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"operating_systems"`
		PercentCalculated int `json:"percent_calculated"`
		Projects          []struct {
			Decimal      string  `json:"decimal"`
			Digital      string  `json:"digital"`
			Hours        int     `json:"hours"`
			Minutes      int     `json:"minutes"`
			Name         string  `json:"name"`
			Percent      float64 `json:"percent"`
			Text         string  `json:"text"`
			TotalSeconds float64 `json:"total_seconds"`
		} `json:"projects"`
		Range                              string    `json:"range"`
		Start                              time.Time `json:"start"`
		Status                             string    `json:"status"`
		Timeout                            int       `json:"timeout"`
		Timezone                           string    `json:"timezone"`
		TotalSeconds                       float64   `json:"total_seconds"`
		TotalSecondsIncludingOtherLanguage float64   `json:"total_seconds_including_other_language"`
		UserID                             string    `json:"user_id"`
		Username                           string    `json:"username"`
		WritesOnly                         bool      `json:"writes_only"`
	} `json:"data"`
}

type languages []struct {
	Decimal      string  `json:"decimal"`
	Digital      string  `json:"digital"`
	Hours        int     `json:"hours"`
	Minutes      int     `json:"minutes"`
	Name         string  `json:"name"`
	Percent      float64 `json:"percent"`
	Text         string  `json:"text"`
	TotalSeconds float64 `json:"total_seconds"`
}
