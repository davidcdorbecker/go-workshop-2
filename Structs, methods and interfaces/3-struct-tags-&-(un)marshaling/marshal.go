package main

import (
	"encoding/json"
	"fmt"
)

// Change the struct to make LastActive appear as last_active in json

type GetUserResponse struct {
	Name       string
	Email      string
	password   string
	Age        int
	LastActive float64
}

func marshal() {
	userResponse := GetUserResponse{
		Name:       "Fer",
		Email:      "idk@idk.com",
		Age:        24,
		password:   "sensibleStuffHere",
		LastActive: 123456,
	}

	resp, err := json.MarshalIndent(userResponse, "", "  ")
	if err != nil {
		fmt.Println("marshal error:", err)
		return
	}

	fmt.Println(string(resp))
}
