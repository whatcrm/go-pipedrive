package utils

const (
	BaseURL                = "https://api.pipedrive.com/v1/"
	DomainBaseURL          = "https://%s.pipedrive.com/"
	WebhooksEndPoint       = "/webhooks?api_token=%s"
	DeleteWebhooksEndPoint = "/webhooks:id?api_token=%s"
	TokenEndPoint          = "https://oauth.pipedrive.com/oauth/token"
	ChannelEndPoint        = "channels"
	DeleteChannelEndPoint  = "channels/%s"
	ReceiveMessageEndPoint = "channels/messages/receive"
	ConversationEndPoint   = "/conversations/"
	UserEndPoint           = "users"
	UserMeEndPoint           = "/me"

)
