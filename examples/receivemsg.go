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
	domain := "<YOUR_COMPANY_DOMAIN>"
	apiToken := "<API_TOKEN>"

	client, err := gopipedrive.NewClient(domain, clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error: ", err)
	}

	client.Token = apiToken

	msg := models.MessageRequest{
		ID:               "1",
		ChannelID:        "5bf042cadb4f2d7d276540579be77e55",
		SenderID:         "77028247806@c.us",
		ConversationID:   "77028247806@c.us",
		Message:          "dddd",
		Status:           "delivered",
		CreatedAt:        "1716403321",
		Attachments: []models.Attachment{
		},
	}

	err = client.ReceiveMessage(context.Background(), msg)

	fmt.Println(" err: ", err)

}
