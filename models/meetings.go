package models

// LinkUserProviderRequest represents the request body for linking a user with a video call provider.
type LinkUserProviderRequest struct {
	UserProviderID       string `json:"user_provider_id"`       // Unique identifier linking a user to the integration
	UserID               int    `json:"user_id"`                // Pipedrive user ID
	CompanyID            int    `json:"company_id"`             // Pipedrive company ID
	MarketplaceClientID  string `json:"marketplace_client_id"`  // Pipedrive Marketplace client ID
}

// LinkUserProviderResponse represents the response from linking a user with a video call provider.
type LinkUserProviderResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// DeleteUserProviderLinkResponse represents the response from deleting a user-provider link.
type DeleteUserProviderLinkResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
