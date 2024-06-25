package main

import (
	"context"
	"fmt"

	gopipedrive "github.com/whatcrm/go-pipedrive"
)

func main() {
	clientID := "<YOUR_CLIENT_ID>"
	clientSecret := "<YOUR_CLIENT_SECRET>"
	redirectURI := "<YOUR_REDIRECT_URI>"
	domain := "<YOUR_COMPANY_DOMAIN>"
	apiToken := "<API_TOKEN>"

	client, err := gopipedrive.NewClient(domain, clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error: ", err)
	}

	client.Token = apiToken

	err = client.DeleteChannel(context.Background(), "5bf042cadb4f2d7d276540579be77e55")

	fmt.Println("err: ", err)

}
