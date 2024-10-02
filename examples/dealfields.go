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

	//Getting all deal fields

	dealFields, err := client.GetDealFields(ctx)
	if err != nil {
		log.Fatal("Error getting deal fields: ", err)
	}

	fmt.Println("dealFields: ", dealFields)

	// Adding new deal field

	newDealField := models.DealFieldReq{
		Name:           "New Label Namme",
		FieldType:      "varchar",
		AddVisibleFlag: true,
	}

	dealField, err := client.AddDealField(ctx, newDealField)
	if err != nil {
		log.Fatal("error adding: ", err)
	}

	fmt.Println("new dealField: ", dealField)

	// Updating deal fields

	newField := models.DealFieldReq{
		Name:           "name",
		AddVisibleFlag: true,
	}

	updatedDealField, err := client.UpdateDealField(ctx, 41, newField)
	if err != nil {
		log.Fatal("Error updating deal field: ", err)
	}

	fmt.Println("updated dealField: ", updatedDealField)

	err = client.DeleteDealField(ctx, 41)
	if err != nil {
		log.Fatal("Error deleting deal field: ", err)
	}

	fmt.Println("deleted succesfully")
}
