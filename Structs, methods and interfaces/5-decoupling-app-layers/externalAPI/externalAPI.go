package externalapi

import "fmt"

type ExternalAPI struct {
}

func (e ExternalAPI) GetInfoFromAPI() {
	fmt.Println("Making a request to external API...")
}
