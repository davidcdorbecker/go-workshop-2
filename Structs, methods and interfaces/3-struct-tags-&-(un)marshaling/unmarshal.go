package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Make the required changes to this video to be able to unmarshal successfully
// the unmarshalDynamic.json file

// part 2, what if we want to save some fields in the DB different from how they appear in json?

type Video struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Views    int      `json:"views"`
	Duration float64  `json:"duration"`
	Likes    int      `json:"likes"`
	Dislikes int      `json:"dislikes"`
	Comments Comments `json:"comments,omitempty"`
}

type Comments []Comment

type Comment struct {
	Author  string
	Content string
}

func unmarshal() {
	normalVideoData, err := os.ReadFile("unmarshal.json")
	if err != nil {
		fmt.Println("error opening file:", err)
		return
	}

	video := &Video{}

	err = json.Unmarshal(normalVideoData, video)
	if err != nil {
		fmt.Println("error unmarhsaling video:", err)
		return
	}

	err = prettyPrint(video)
	if err != nil {
		fmt.Println("error printing video data:", err)
		return
	}
}

func prettyPrint(v *Video) error {
	videoData, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}

	fmt.Println(string(videoData))
	return nil
}
