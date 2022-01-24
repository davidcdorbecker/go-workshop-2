package db

import "fmt"

type DB struct {
}

func (d DB) GetInfoFromDB() {
	fmt.Println("Querying DB...")
}
