package analyzer

import (
	"fmt"
	"log-parser/pkg/parser"
	"time"
)

type Analyzer struct {
	filterIP  string
	startDate string
	endDate   string
}

func (a Analyzer) CountStatusCodes(stats []parser.LogStat) map[int]int {
	statusCodes := make(map[int]int)
	for _, record := range stats {
		if a.shouldSkipRecord(record) {
			continue
		}
		if a.shouldBreak(record) {
			break
		}
		statusCodes[record.Status]++
	}
	return nil
}

func (a Analyzer) CountUniqueIPs(stats []parser.LogStat) map[string]int {
	uniqueIPs := make(map[string]int)
	for _, record := range stats {
		if a.shouldSkipRecord(record) {
			continue
		}
		if a.shouldBreak(record) {
			break
		}
		uniqueIPs[record.IP]++
	}
	return uniqueIPs
}

func (a Analyzer) AveregeBytes(stats []parser.LogStat) float64 {
	var totalBytes int
	for _, record := range stats {
		if a.shouldSkipRecord(record) {
			continue
		}
		if a.shouldBreak(record) {
			break
		}
		totalBytes += record.BytesSent
	}
	return float64(totalBytes) / float64(len(stats))
}

func (a Analyzer) TopUrls(stats []parser.LogStat) map[string]int {
	urls := make(map[string]int)
	topUrls := make(map[string]int)
	const minToTop = 5
	for _, record := range stats {
		if a.shouldSkipRecord(record) {
			continue
		}
		if a.shouldBreak(record) {
			break
		}
		urls[record.Path]++
		if urls[record.Path] >= minToTop {
			topUrls[record.Path] = urls[record.Path]
		}
	}
	return topUrls
}

func (a Analyzer) shouldSkipRecord(record parser.LogStat) bool {
	if a.filterIP != "" && record.IP != a.filterIP {
		return true
	}
	startDate, err := time.Parse("02/Jan/2006:15:04:05 -0700", a.startDate)
	if err != nil {
		fmt.Println("Error parsing start date: ", err)
	}

	// Parse record.TimeStamp
	recordTime, err := time.Parse("02/Jan/2006:15:04:05 -0700", record.TimeStamp)
	if err != nil {
		fmt.Println("Error parsing record time: ", err)
	}
	if recordTime.Before(startDate) {
		return true
	}
	return false
}

func (a Analyzer) shouldBreak(record parser.LogStat) bool {
	endDate, err := time.Parse("02/Jan/2006:15:04:05 -0700", a.endDate)
	if err != nil {
		fmt.Println("Error parsing end date: ", err)
		return false
	}
	recordTime, err := time.Parse("02/Jan/2006:15:04:05 -0700", record.TimeStamp)
	if err != nil {
		fmt.Println("Error parsing record time: ", err)
		return false
	}
	if recordTime.After(endDate) {
		return true
	}
	return false
}
