package main

import (
	"context"
	"fmt"

	gopipedrive "github.com/whatcrm/go-pipedrive"
)

func main() {
	clientID := "<YOUR_CLIENT_ID>"
	clientSecret := "<YOUR_CLIENT_SECRET>"

	authorizationCode := "<YOUR_AUTH_CODE>"
	redirectURI := "<YOUR_REDIRECT_URI>"

	client, err := gopipedrive.NewClient(clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error: ", err)
	}

	ctx := context.Background()

	token, err := client.GetAccessToken(ctx, authorizationCode)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("result: ", token)
}
