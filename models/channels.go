package models

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
 