package gopipedrive

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/whatcrm/go-pipedrive/models"
	"github.com/whatcrm/go-pipedrive/utils"
)

// GetAllDeals retrieves all deals with optional query parameters
func (c *Client) GetAllDeals(ctx context.Context, queryParams map[string]string) ([]models.Deal, error) {
	url := c.APIBase + utils.DealsEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dealsResponse models.DealsResponse
	err = c.SendWithAccessToken(req, &dealsResponse)
	if err != nil {
		return nil, err
	}

	return dealsResponse.Data, nil
}

// GetDeal retrieves a specific deal by ID
func (c *Client) GetDeal(ctx context.Context, dealID int) (models.Deal, error) {
	dealEndpoint := fmt.Sprintf(utils.OneDealEndPoint, dealID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.Deal{}, err
	}

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// AddDeal creates a new deal
func (c *Client) AddDeal(ctx context.Context, dealReq models.DealRequest) (models.Deal, error) {
	url := c.APIBase + utils.DealsEndpoint

	requestBodyBytes, err := json.Marshal(dealReq)
	if err != nil {
		return models.Deal{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Deal{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// UpdateDeal updates an existing deal
func (c *Client) UpdateDeal(ctx context.Context, dealID int, dealUpdateReq models.DealUpdateRequest) (models.Deal, error) {
	dealEndpoint := fmt.Sprintf(utils.OneDealEndPoint, dealID)
	url := c.APIBase + dealEndpoint

	requestBodyBytes, err := json.Marshal(dealUpdateReq)
	if err != nil {
		return models.Deal{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Deal{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// DeleteDeal deletes a deal
func (c *Client) DeleteDeal(ctx context.Context, dealID int) error {
	dealEndpoint := fmt.Sprintf(utils.OneDealEndPoint, dealID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// AddDealFollower adds a follower to a deal
func (c *Client) AddDealFollower(ctx context.Context, dealID, userID int) (models.DealFollower, error) {
	dealEndpoint := fmt.Sprintf(utils.DealFollowersEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	followerReq := models.DealFollowerRequest{UserID: userID}
	requestBodyBytes, err := json.Marshal(followerReq)
	if err != nil {
		return models.DealFollower{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.DealFollower{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var followerResponse models.DealFollowerResponse
	err = c.SendWithAccessToken(req, &followerResponse)
	if err != nil {
		return models.DealFollower{}, err
	}

	return followerResponse.Data, nil
}

// DeleteDealFollower deletes a follower from a deal
func (c *Client) DeleteDealFollower(ctx context.Context, dealID, followerID int) error {
	dealEndpoint := fmt.Sprintf(utils.DealOneFollowerEndpoint, dealID, followerID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// AddDealParticipant adds a participant to a deal
func (c *Client) AddDealParticipant(ctx context.Context, dealID, personID int) (models.DealParticipant, error) {
	dealEndpoint := fmt.Sprintf(utils.DealParticipantsEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	participantReq := models.DealParticipantRequest{PersonID: personID}
	requestBodyBytes, err := json.Marshal(participantReq)
	if err != nil {
		return models.DealParticipant{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.DealParticipant{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var participantResponse models.DealParticipantResponse
	err = c.SendWithAccessToken(req, &participantResponse)
	if err != nil {
		return models.DealParticipant{}, err
	}

	return participantResponse.Data, nil
}

// DeleteDealParticipant deletes a participant from a deal
func (c *Client) DeleteDealParticipant(ctx context.Context, dealID, participantID int) error {
	dealEndpoint := fmt.Sprintf(utils.DealOneParticipantEndpoint, dealID, participantID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// AddDealProduct adds a product to a deal
func (c *Client) AddDealProduct(ctx context.Context, dealID int, productReq models.DealProductRequest) (models.DealProduct, error) {
	dealEndpoint := fmt.Sprintf(utils.DealProductsEndpoint, dealID)
	url := c.APIBase + dealEndpoint

	requestBodyBytes, err := json.Marshal(productReq)
	if err != nil {
		return models.DealProduct{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.DealProduct{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productResponse models.DealProductResponse
	err = c.SendWithAccessToken(req, &productResponse)
	if err != nil {
		return models.DealProduct{}, err
	}

	return productResponse.Data, nil
}

// UpdateDealProduct updates the details of a product attached to a deal
func (c *Client) UpdateDealProduct(ctx context.Context, dealID, productAttachmentID int, productUpdateReq models.DealProductUpdateRequest) (models.DealProduct, error) {
	dealEndpoint := fmt.Sprintf(utils.DealOneProductEndpoint, dealID, productAttachmentID)
	url := c.APIBase + dealEndpoint

	requestBodyBytes, err := json.Marshal(productUpdateReq)
	if err != nil {
		return models.DealProduct{}, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.DealProduct{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productResponse models.DealProductResponse
	err = c.SendWithAccessToken(req, &productResponse)
	if err != nil {
		return models.DealProduct{}, err
	}

	return productResponse.Data, nil
}

// DeleteDealProduct deletes a product attachment from a deal
func (c *Client) DeleteDealProduct(ctx context.Context, dealID, productAttachmentID int) error {
	dealEndpoint := fmt.Sprintf(utils.DealOneProductEndpoint, dealID, productAttachmentID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// AddDealDiscount adds a discount to a deal
func (c *Client) AddDealDiscount(ctx context.Context, dealID int, discountReq models.DealDiscountRequest) (models.DealDiscount, error) {
	dealEndpoint := fmt.Sprintf(utils.DealDiscountsEndpoint, dealID)
	url := c.APIBase + dealEndpoint

	requestBodyBytes, err := json.Marshal(discountReq)
	if err != nil {
		return models.DealDiscount{}, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.DealDiscount{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var discountResponse models.DealDiscountResponse
	err = c.SendWithAccessToken(req, &discountResponse)
	if err != nil {
		return models.DealDiscount{}, err
	}

	return discountResponse.Data, nil
}

// UpdateDealDiscount updates a discount attached to a deal
func (c *Client) UpdateDealDiscount(ctx context.Context, dealID, discountID int, discountUpdateReq models.DealDiscountUpdateRequest) (models.DealDiscount, error) {
	dealEndpoint := fmt.Sprintf(utils.DealOneDiscountEndpoint, dealID, discountID)
	url := c.APIBase + dealEndpoint

	requestBodyBytes, err := json.Marshal(discountUpdateReq)
	if err != nil {
		return models.DealDiscount{}, err
	}

	req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.DealDiscount{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var discountResponse models.DealDiscountResponse
	err = c.SendWithAccessToken(req, &discountResponse)
	if err != nil {
		return models.DealDiscount{}, err
	}

	return discountResponse.Data, nil
}

// DeleteDealDiscount deletes a discount from a deal
func (c *Client) DeleteDealDiscount(ctx context.Context, dealID, discountID int) error {
	dealEndpoint := fmt.Sprintf(utils.DealOneDiscountEndpoint, dealID, discountID)
	url := c.APIBase + dealEndpoint
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	return c.SendWithAccessToken(req, nil)
}

// MergeDeals merges two deals
func (c *Client) MergeDeals(ctx context.Context, dealID, mergeWithID int) (models.Deal, error) {
	dealEndpoint := fmt.Sprintf(utils.DealsMergeEndpoint, dealID)
	url := c.APIBase + dealEndpoint

	mergeReq := models.MergeDealRequest{MergeWithID: mergeWithID}
	requestBodyBytes, err := json.Marshal(mergeReq)
	if err != nil {
		return models.Deal{}, err
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(requestBodyBytes))
	if err != nil {
		return models.Deal{}, err
	}
	req.Header.Set("Content-Type", "application/json")

	var dealResponse models.DealResponse
	err = c.SendWithAccessToken(req, &dealResponse)
	if err != nil {
		return models.Deal{}, err
	}

	return dealResponse.Data, nil
}

// GetAllDealsBeta retrieves all deals using the BETA endpoint
func (c *Client) GetAllDealsBeta(ctx context.Context, queryParams map[string]string) ([]models.Deal, error) {
	url := c.APIBase + utils.DealsCollectionEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var dealsResponse models.DealsResponse
	err = c.SendWithAccessToken(req, &dealsResponse)
	if err != nil {
		return nil, err
	}

	return dealsResponse.Data, nil
}

// SearchDeals searches for deals
func (c *Client) SearchDeals(ctx context.Context, queryParams map[string]string) ([]models.Deal, error) {
	url := utils.SearchDealsURL
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var searchResponse models.DealsResponse
	err = c.SendWithAccessToken(req, &searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse.Data, nil
}

// GetDealsSummary retrieves a summary of all deals
func (c *Client) GetDealsSummary(ctx context.Context, queryParams map[string]string) (models.DealsSummaryResponse, error) {
	url := c.APIBase + utils.DealsSummaryEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.DealsSummaryResponse{}, err
	}

	var summaryResponse models.DealsSummaryResponse
	err = c.SendWithAccessToken(req, &summaryResponse)
	if err != nil {
		return models.DealsSummaryResponse{}, err
	}

	return summaryResponse, nil
}

// GetDealsTimeline retrieves open and won deals grouped by a defined interval
func (c *Client) GetDealsTimeline(ctx context.Context, queryParams map[string]string) (models.DealsTimelineResponse, error) {
	url := c.APIBase + utils.DealsTimelineEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return models.DealsTimelineResponse{}, err
	}

	var timelineResponse models.DealsTimelineResponse
	err = c.SendWithAccessToken(req, &timelineResponse)
	if err != nil {
		return models.DealsTimelineResponse{}, err
	}

	return timelineResponse, nil
}

// ListDealActivities lists activities associated with a deal
func (c *Client) ListDealActivities(ctx context.Context, dealID int, queryParams map[string]string) ([]models.Activity, error) {
	dealEndpoint := fmt.Sprintf(utils.DealActivitiesEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var activitiesResponse models.ActivitiesResponse
	err = c.SendWithAccessToken(req, &activitiesResponse)
	if err != nil {
		return nil, err
	}

	return activitiesResponse.Data, nil
}

// ListDealFieldUpdates lists updates about field values of a deal
func (c *Client) ListDealFieldUpdates(ctx context.Context, dealID int, queryParams map[string]string) ([]models.DealFieldUpdate, error) {
	dealEndpoint := fmt.Sprintf(utils.DealFieldUpdatesEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var fieldUpdatesResponse models.DealFieldUpdatesResponse
	err = c.SendWithAccessToken(req, &fieldUpdatesResponse)
	if err != nil {
		return nil, err
	}

	return fieldUpdatesResponse.Data, nil
}

// ListDealFiles lists files attached to a deal
func (c *Client) ListDealFiles(ctx context.Context, dealID int, queryParams map[string]string) ([]models.File, error) {
	dealEndpoint := fmt.Sprintf(utils.DealFilesEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var filesResponse models.FilesResponse
	err = c.SendWithAccessToken(req, &filesResponse)
	if err != nil {
		return nil, err
	}

	return filesResponse.Data, nil
}

// ListDealUpdates lists updates about a deal
func (c *Client) ListDealUpdates(ctx context.Context, dealID int, queryParams map[string]string) ([]models.DealUpdate, error) {
	dealEndpoint := fmt.Sprintf(utils.DealUpdatesEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var updatesResponse models.DealUpdatesResponse
	err = c.SendWithAccessToken(req, &updatesResponse)
	if err != nil {
		return nil, err
	}

	return updatesResponse.Data, nil
}

// ListDealParticipants lists participants associated with a deal
func (c *Client) ListDealParticipants(ctx context.Context, dealID int, queryParams map[string]string) ([]models.Participant, error) {
	dealEndpoint := fmt.Sprintf(utils.DealParticipantsEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var participantsResponse models.ParticipantsResponse
	err = c.SendWithAccessToken(req, &participantsResponse)
	if err != nil {
		return nil, err
	}

	return participantsResponse.Data, nil
}

// ListDealFollowers lists followers of a deal
func (c *Client) ListDealFollowers(ctx context.Context, dealID int) ([]models.Follower, error) {
	dealEndpoint := fmt.Sprintf(utils.DealFollowersEndpoint, dealID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var followersResponse models.FollowersResponse
	err = c.SendWithAccessToken(req, &followersResponse)
	if err != nil {
		return nil, err
	}

	return followersResponse.Data, nil
}

// ListDealMailMessages lists mail messages associated with a deal
func (c *Client) ListDealMailMessages(ctx context.Context, dealID int, queryParams map[string]string) ([]models.MailMessage, error) {
	dealEndpoint := fmt.Sprintf(utils.DealMailMessagesEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var mailMessagesResponse models.MailMessagesResponse
	err = c.SendWithAccessToken(req, &mailMessagesResponse)
	if err != nil {
		return nil, err
	}

	return mailMessagesResponse.Data, nil
}

// ListDealPermittedUsers lists the users permitted to access a deal
func (c *Client) ListDealPermittedUsers(ctx context.Context, dealID int) ([]models.User, error) {
	dealEndpoint := fmt.Sprintf(utils.DealPermittedUsersEndpoint, dealID)
	url := c.APIBase + dealEndpoint

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var usersResponse models.UsersResponse
	err = c.SendWithAccessToken(req, &usersResponse)
	if err != nil {
		return nil, err
	}

	return usersResponse.Data, nil
}

// ListDealPersons lists all persons associated with a deal
func (c *Client) ListDealPersons(ctx context.Context, dealID int, queryParams map[string]string) ([]models.Person, error) {
	dealEndpoint := fmt.Sprintf(utils.DealPersonsEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var personsResponse models.PersonsResponse
	err = c.SendWithAccessToken(req, &personsResponse)
	if err != nil {
		return nil, err
	}

	return personsResponse.Data, nil
}

// ListDealProducts lists products attached to a deal
func (c *Client) ListDealProducts(ctx context.Context, dealID int, queryParams map[string]string) ([]models.Product, error) {
	dealEndpoint := fmt.Sprintf(utils.DealProductsEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var productsResponse models.ProductsResponse
	err = c.SendWithAccessToken(req, &productsResponse)
	if err != nil {
		return nil, err
	}

	return productsResponse.Data, nil
}

// GetDealProductsOfSeveralDeals retrieves products attached to several deals
func (c *Client) GetDealProductsOfSeveralDeals(ctx context.Context, dealIDs []int, queryParams map[string]string) ([]models.Product, error) {
	url := utils.DealsProductsURL
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	reqBody, err := json.Marshal(map[string]interface{}{
		"deal_ids": dealIDs,
	})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	var productsResponse models.ProductsResponse
	err = c.SendWithAccessToken(req, &productsResponse)
	if err != nil {
		return nil, err
	}

	return productsResponse.Data, nil
}

// ListDealDiscounts lists discounts attached to a deal
func (c *Client) ListDealDiscounts(ctx context.Context, dealID int, queryParams map[string]string) ([]models.Discount, error) {
	dealEndpoint := fmt.Sprintf(utils.DealDiscountsEndpoint, dealID)
	url := c.APIBase + dealEndpoint
	if len(queryParams) > 0 {
		url += "?"
		for key, value := range queryParams {
			url += fmt.Sprintf("%s=%s&", key, value)
		}
		url = url[:len(url)-1]
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var discountsResponse models.DiscountsResponse
	err = c.SendWithAccessToken(req, &discountsResponse)
	if err != nil {
		return nil, err
	}

	return discountsResponse.Data, nil
}
