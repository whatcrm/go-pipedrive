package models

import "time"

type ChannelRequest struct {
	Name              string `json:"name"`
	ProviderChannelID string `json:"provider_channel_id"`
	AvatarURL         string `json:"avatar_url,omitempty"`
	TemplateSupport   bool   `json:"template_support,omitempty"`
	ProviderType      string `json:"provider_type,omitempty"`
}

type MessageRequest struct {
	ID               string   `json:"id"`
	ChannelID        string   `json:"channel_id"`
	SenderID         string   `json:"sender_id"`
	ConversationID   string   `json:"conversation_id"`
	Message          string   `json:"message"`
	Status           string   `json:"status"`
	CreatedAt        string   `json:"created_at"`
	ReplyBy          string   `json:"reply_by,omitempty"`
	ConversationLink string   `json:"conversation_link,omitempty"`
	Attachments      []string `json:"attachments,omitempty"`
}

type ChannelReponse struct {
	Success bool    `json:"success"`
	Data    Channel `json:"data"`
}

type Channel struct {
	ID                  string    `json:"id"`
	Name                string    `json:"name"`
	AvatarURL           string    `json:"avatar_url"`
	ProviderChannelID   string    `json:"provider_channel_id"`
	MarketplaceClientID string    `json:"marketplace_client_id"`
	PDCompanyID         int       `json:"pd_company_id"`
	PDUserID            int       `json:"pd_user_id"`
	CreatedAt           time.Time `json:"created_at"`
	ProviderType        string    `json:"provider_type"`
	TemplateSupport     bool      `json:"template_support"`
}
