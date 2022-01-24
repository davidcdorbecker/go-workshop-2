package main

import "fmt"

type Viewer interface {
	WatchVideo(videoName string)
}

type Commenter interface {
	AddComment(commentContent string)
}

type User struct {
	Name string
}

func (u User) WatchVideo(videoName string) {
	fmt.Printf("User %s is watching %s video!\n", u.Name, videoName)
}

func (u User) AddComment(commentContent string) {
	fmt.Printf("Comment to add: %s\n", commentContent)

}

func composite() {
	user := User{
		Name: "Fer",
	}

	watchAndCommentVideo(user, user)
}

// Find a way to improve the usage of this function and receive only one parameter
// that can do both the Viewer and the Commenter job
func watchAndCommentVideo(viewer Viewer, commenter Commenter) {
	viewer.WatchVideo("In memory of Harambe")
	commenter.AddComment("I give this video a perfect 5/7")
}
