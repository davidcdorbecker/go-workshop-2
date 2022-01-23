package analytics

import "fmt"

type VideoAnalytics struct {
	AverageViewTime float64
	MainCountry     string
	Ads             []string
}

func (va VideoAnalytics) GetAnalyticsFormatted() string {
	return fmt.Sprintf("Average View Time: %f\nMainCountry: %s\nAds: %v\n", va.AverageViewTime, va.MainCountry, va.Ads)
}
