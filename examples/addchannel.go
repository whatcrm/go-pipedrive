package main

import (
	"context"
	"fmt"

	gopipedrive "github.com/whatcrm/go-pipedrive"
	"github.com/whatcrm/go-pipedrive/models"
)

func main() {
	clientID := "<YOUR_CLIENT_ID>"
	clientSecret := "<YOUR_CLIENT_SECRET>"
	redirectURI := "<YOUR_REDIRECT_URI>"
	apiToken := "<API_TOKEN>"

	client, err := gopipedrive.NewClient(clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error: ", err)
	}

	client.Token = apiToken

	channel := models.ChannelRequest{
		Name:              "",
		ProviderChannelID: "",
		TemplateSupport:   false,
		ProviderType:      "whatsapp",
	}

	c, err := client.AddChannel(context.Background(), channel)

	fmt.Println("res: ", c, " err: ", err)

}
