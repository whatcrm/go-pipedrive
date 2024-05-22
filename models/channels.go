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
	ID               string       `json:"id"`
	SenderID         string       `json:"sender_id"`
	ConversationID   string       `json:"conversation_id"`
	ConversationLink *string      `json:"conversation_link,omitempty"`
	ChannelID        string       `json:"channel_id"`
	CreatedAt        string       `json:"created_at"`
	Message          string       `json:"message"`
	Status           interface{}  `json:"status"`
	Attachments      []Attachment `json:"attachments"`
	ReplyBy          *string      `json:"reply_by,omitempty"`
}

type Attachment struct {
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Name        *string `json:"name"`
	Size        *int    `json:"size"`
	URL         string  `json:"url"`
	PreviewURL  *string `json:"preview_url"`
	LinkExpires *bool   `json:"link_expires,omitempty"`
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
