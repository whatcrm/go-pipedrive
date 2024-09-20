package main

import (
	"context"
	"fmt"
	"log"

	gopipedrive "github.com/whatcrm/go-pipedrive"
	"github.com/whatcrm/go-pipedrive/models"
)

type Option struct {
	Label string `json:"label"`
}

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

	newLabelField := models.PersonFieldReq{
		Name:      "Label",   // Name of the new label
		FieldType: "enum",        // Setting the type to 'enum' for a dropdown option
		Options: []Option{ // Define the options for the new label field
			{Label: "VIP"},
			{Label: "Regular"},
			{Label: "New Customer"},
		},
		AddVisibleFlag: true, // Make it visible in the 'Add New' modal
	}

	personField, err := client.AddPersonField(context.Background(), newLabelField)
	if err != nil {
		log.Fatal("Error adding new label to person: ", err)
	}

	fmt.Println("New label added to person field:", personField)

	personFields, err := client.GetPersonFields(context.Background())
	if err != nil {
		log.Fatal("Error getting person fields: ", err)
	}

	fmt.Println("Person fields:")
	for _, field := range *personFields {
		fmt.Printf("ID: %d, Name: %s, Type: %s\n", field.ID, field.Name, field.FieldType)
	}
}
