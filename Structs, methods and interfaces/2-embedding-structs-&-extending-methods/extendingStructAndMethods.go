package main

import (
	"fmt"

	"embeddingAndExtending/analytics"
)

/*
	Given that there's no inheritance on Go or extension methods, how can we add a new method to the VideoAnalytics structure
	on this file? Maybe also add new fields?
*/

func extendingStructsAndMethods() {
	a := analytics.VideoAnalytics{
		AverageViewTime: 5.5,
		MainCountry:     "USA",
		Ads:             []string{"Keeps", "Nord VPN"},
	}

	fmt.Println(a.GetAnalyticsFormatted())
}
