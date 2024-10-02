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
		log.Fatal("error: ", err)
	}

	client.Token = apiToken
	ctx := context.Background()

	//Getting all person fields

	personFields, err := client.GetPersonFields(ctx)
	if err != nil {
		log.Fatal("Error getting person fields: ", err)
	}

	fmt.Println("personFields: ", personFields)

	// Adding new person field

	newPersonField := models.PersonFieldReq{
		Name:           "New Label Namme",
		FieldType:      "varchar",
		AddVisibleFlag: true,
	}

	personField, err := client.AddPersonField(ctx, newPersonField)
	if err != nil {
		log.Fatal("error adding: ", err)
	}

	fmt.Println("new personField: ", personField)

	// Updating person fiels (adding opton to Labels)
	var existingOptions []models.Option
	var personFieldID int

	for _, p := range *personFields {
		if p.Name == "Labels" {
			existingOptions = p.Options
			personFieldID = p.ID
		}
	}

	newField := models.PersonFieldReq{
		Options:        append(existingOptions, models.Option{Label: "New test label"}),
		AddVisibleFlag: true,
	}

	updatedPersonField, err := client.UpdatePersonField(ctx, personFieldID, newField)
	if err != nil {
		log.Fatal("Error updating person field: ", err)
	}

	fmt.Println("updated personField: ", updatedPersonField)

	// err = client.DeletePersonField(ctx, 28)
	// if err != nil {
	// 	log.Fatal("Error deleting person field: ", err)
	// }

	fmt.Println("deleted succesfully")

}
