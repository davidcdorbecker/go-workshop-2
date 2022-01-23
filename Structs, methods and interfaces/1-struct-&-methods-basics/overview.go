package main

import "fmt"

type Video struct {
	ID       string
	Name     string
	Duration float64
	Views    int
	Likes    int
	dislikes int
	Comments []string
}

func overview() {
	video1 := Video{
		"07d0a37a-f4f2-41cc-9f05-5fb17a7a6dfb",
		"Why has Go taken so long to implement generics?",
		10.30,
		1000000,
		10000,
		500,
		[]string{"yeah, totally", "they're finally coming for Go 1.18 *tears of happiness*"},
	}

	video2 := &Video{
		ID:       "7ac61580-75b4-451c-81bd-f9c620f8a793",
		Name:     "Is tail recursion supported on Go?",
		Comments: []string{"nah"},
	}

	// wrongVideo := Video{}

	changeVideoStats(video1)
	fmt.Printf("\nvideo 1 after changing video stats: %+v\n\n\n", video1)

	changeVideoStatsPointer(video2)
	fmt.Printf("\nvideo 2 after changing video stats: %+v\n", video2)
}

// Try to change the attributes of a video and see what happens for certain data types
func changeVideoStats(v Video) {

}

// Try to change the attributes of a pointer to a video and see what happens
func changeVideoStatsPointer(v *Video) {

}
