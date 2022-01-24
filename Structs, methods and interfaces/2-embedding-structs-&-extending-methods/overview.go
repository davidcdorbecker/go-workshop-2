package main

import "fmt"

type Video struct {
	ID       string
	Author   string
	Name     string
	Duration float64
	Views    int
	Likes    int
	dislikes int
}

func (v Video) getAuthor() string {
	return v.Author
}

// Make a shortVideo behave and have the same properties and behave in the same way a normal Video would
type ShortVideo struct {
	ID       string
	Lifetime float64
	Video    Video
}

func overview() {
	sv := ShortVideo{
		Video: Video{
			Author: "Fer",
		},
	}

	fmt.Println(sv.Video.getAuthor())
}
