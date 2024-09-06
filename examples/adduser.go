package main

import (
	"context"
	"fmt"
	"log"

	gopipedrive "github.com/whatcrm/go-pipedrive"
	"github.com/whatcrm/go-pipedrive/models"
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

	newUser := models.CreateUserReq{
		Email: "email@mail.com",
		Access: []models.Access{
			{App: "sales"},
		},
		ActiveFlag: true,
	}

	user, err := client.AddUser(context.Background(), newUser)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(user)

}
