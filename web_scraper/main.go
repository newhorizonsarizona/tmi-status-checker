package main

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
	"gopkg.in/yaml.v3"
)

func main() {
	c := colly.NewCollector()
	// Get the current year
	currentTime := time.Now()
	currentMonth := currentTime.Month()
	currentYear := currentTime.Year()
	nextYear := currentYear + 1
	if currentMonth >= 1 && currentMonth <= 6 {
		currentYear = currentYear - 1
		nextYear = nextYear - 1
	}
	tmYear := strconv.Itoa(currentYear) + "-" + strconv.Itoa(nextYear)
	clubReportUrl := "https://dashboards.toastmasters.org/ClubReport.aspx?id=00006350"
	dcpGoals := make(map[string]map[string]map[string]string)
	// On every <a> element which has href attribute call callback
	goalsFound := false
	memberSummaryFound := false
	memberSummaryBaseFound := false
	membershipBase := ""
	memberSummaryToDateFound := false
	membershipToDate := ""
	goalCategoryKey := ""
	goalKey := ""
	goalValue := map[string]string{"target": "", "achieved": "", "status": ""}
	dcpStatusRef := map[int]string{
		5:  "Distinguished",
		6:  "Distinguished",
		7:  "Select Distinguished",
		8:  "Select Distinguished",
		9:  "President's Distinguished",
		10: "President's Distinguished",
	}
	currentAchievementCount := 0
	c.OnHTML("table th", func(e *colly.HTMLElement) {
		if !memberSummaryFound && e.Text != "Membership" {
			memberSummaryFound = true
			return
		}
	})
	c.OnHTML("table td", func(e *colly.HTMLElement) {
		class := e.Attr("class")
		if !goalsFound && e.Text != "Goals to Achieve" {
			if memberSummaryFound {
				if !memberSummaryBaseFound && class == "chart_table_content" {
					memberSummaryBaseFound = true
					return
				}
				if memberSummaryBaseFound && membershipBase == "" && class == "chart_table_big_numbers" {
					membershipBase = e.Text
					return
				}
				if !memberSummaryToDateFound && class == "chart_table_content" {
					memberSummaryToDateFound = true
					return
				}
				if memberSummaryToDateFound && membershipToDate == "" && class == "chart_table_big_numbers" {
					membershipToDate = e.Text
					return
				}
			}
			return
		}
		goalsFound = true
		if class == "categorySeparator" && e.Text != "" {
			goalCategoryKey = e.Text
			dcpGoals[goalCategoryKey] = map[string]map[string]string{}
			goalKey = ""
			return
		}
		if class == "goalDescription" {
			goalKey = strings.TrimSpace(strings.ReplaceAll(e.Text, "All Pathways education awards must be submitted in both Base Camp and Club Central.", ""))
			goalKey = strings.ReplaceAll(goalKey, ",", ";")
			goalValue = map[string]string{"target": "", "achieved": "", "status": ""}
			dcpGoals[goalCategoryKey][goalKey] = goalValue
			return
		}
		if class == "clubReportGoalText" {
			goalValue["target"] = e.Text
			return
		}
		if strings.Contains(class, "clubReportGoalTextAchieved") || strings.Contains(class, "clubReportGoal") {
			goalValue["achieved"] = e.Text
			return
		}
		if class == "statusImage" {
			status := e.Text
			if status == "" {
				e.ForEach("img", func(_ int, el *colly.HTMLElement) {
					imgSrc := el.Attr("src")
					if strings.Contains(imgSrc, "checkMark") {
						status = "Achieved"
					} else {
						status = "Not Achieved"
					}
				})
			}
			goalValue["status"] = status
			if status == "Achieved" {
				currentAchievementCount = currentAchievementCount + 1
			}
		}
		if goalKey == "" {
			return
		}
		dcpGoals[goalCategoryKey][goalKey] = goalValue
	})

	// On request error
	c.OnError(func(r *colly.Response, err error) {
		log.Fatalf("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit(clubReportUrl)
	currentStatus := ""
	for goalTarget, statusName := range dcpStatusRef {
		if currentAchievementCount >= goalTarget {
			currentStatus = statusName
		}
	}
	dcpReport := make(map[string]map[string]map[string]map[string]string)
	dcpGoals["DCP Status"] = map[string]map[string]string{
		"Overall":    {"Year": tmYear, "Distinguished": "No", "Current": strconv.Itoa(currentAchievementCount), "Target": strconv.Itoa(10)},
		"Membership": {"Base": membershipBase, "To Date": membershipToDate, "Required": "20"},
	}
	if currentStatus != "" {
		if currentStatus != "Distinguished" {
			delete(dcpGoals["DCP Status"]["Overall"], "Distinguished")
		}
		dcpGoals["DCP Status"]["Overall"][currentStatus] = "Yes"
	}
	dcpReport["DCP Report"] = dcpGoals
	yamlBytes, err := yaml.Marshal(dcpReport)
	if err != nil {
		log.Fatalf("Error converting map to yml", err)
	}

	// Write the YAML data to a file
	fileName := "../reports/dcp_report.yaml"
	err = os.WriteFile(fileName, yamlBytes, 0644)
	if err != nil {
		log.Fatalf("Failed to write DCP Achievements Yaml file: %v", err)
	}
}
