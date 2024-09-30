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

	msg := models.MessageRequest{
		ID:               "1",
		ChannelID:        "channelid",
		SenderID:         "",
		ConversationID:   "",
		Message:          "",
		Status:           "delivered",
		CreatedAt:        "1716403321",
		Attachments: []models.Attachment{
		},
	}

	err = client.ReceiveMessage(context.Background(), msg)

	fmt.Println(" err: ", err)

}
