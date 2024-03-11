package utils

const (
	BaseURL                = "https://api.pipedrive.com/v1"
	WebhooksEndPoint       = "/webhooks?api_token=%s"
	DeleteWebhooksEndPoint = "/webhooks:id?api_token==%s"
	OAuthURLEndPoint       = "https://oauth.pipedrive.com/oauth"
	AuthorizeEndPoint      = "authorize?client_id=%s&redirect_uri=%s&state=%s"
	TokenEndPoint          = "token"
)
