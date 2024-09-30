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

	// personFields, err := client.GetPersonFields(ctx)
	// if err != nil {
	// 	log.Fatal("Error getting person fields: ", err)
	// }

	// var existingOptions []models.Option
	// var personFieldID int

	// for _, p := range *personFields {
	// 	if p.Name == "Label" {
	// 		existingOptions = p.Options
	// 		personFieldID = p.ID
	// 	}
	// }

	// newField := models.PersonFieldReq{
	// 	Name: "Label",
	// 	Options: append(existingOptions, models.Option{Label: "New test label"}),
	// 	AddVisibleFlag: true,
	// }

	// personField, err := client.UpdatePersonField(ctx, personFieldID, newField)
	// if err != nil {
	// 	log.Fatal("Error updating person field: ", err)
	// }

	// fmt.Print(personField)

	newPersonField := models.PersonFieldReq{
		Name: "Label By Whatcrm",
		FieldType: "varchar",
		AddVisibleFlag: true,
	}

	persoField, err := client.AddPersonField(ctx, newPersonField)
	if err != nil {
		log.Fatal("error adding: ", err)
	}

	fmt.Print(persoField)

}
