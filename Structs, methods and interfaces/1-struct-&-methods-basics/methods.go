package main

import "fmt"

func methods() {
	video1 := Video{
		Likes:    50,
		dislikes: 10,
		Views:    1000,
	}

	fmt.Println(video1.getLikesAndDislikes())

	video1.ChangeViews(5)
	fmt.Printf("\nVideo after changing views: %+v\n\n", video1)

	// What happens if we try to use a nil video to call methods?
}

func (v Video) getLikesAndDislikes() string {
	v.Likes = 10
	return fmt.Sprintf("Likes: %d \nDislikes: %d", v.Likes, v.dislikes)
}

func (v *Video) ChangeViews(views int) {
	v.Views = views
}

// Create a new function with a pointer receiver to modify likes and dislikes
