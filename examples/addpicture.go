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
	apiToken := "<API_TOKEN>"

	client, err := gopipedrive.NewClient(clientID, clientSecret, redirectURI)
	if err != nil {
		fmt.Println("error:", err)
	}

	client.Token = apiToken
	ctx := context.Background()

	imageURL := "<IMAGE_URL>"
	personID := 186

	req := models.PersonPictureRequest{
		FilePath: imageURL,
	}

	personPictureResponse, err := client.AddPersonPicture(ctx, personID, req)
	if err != nil {
		log.Fatalf("Failed to upload picture: %v", err)
	}

	fmt.Printf("Picture uploaded successfully!\n")
	fmt.Printf("Person ID: %d\n", personPictureResponse.ItemID)
	fmt.Printf("128px image URL: %s\n", personPictureResponse.Images.Image128)
	fmt.Printf("512px image URL: %s\n", personPictureResponse.Images.Image512)
}
