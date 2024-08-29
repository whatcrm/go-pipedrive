package models

type Webhook struct {
	ID               int     `json:"id"`
	CompanyID        int     `json:"company_id"`
	OwnerID          int     `json:"owner_id"`
	UserID           int     `json:"user_id,omitempty"`
	EventAction      string  `json:"event_action"`
	EventObject      string  `json:"event_object"`
	SubscriptionURL  string  `json:"subscription_url"`
	Version          string  `json:"version,omitempty"`
	IsActive         int     `json:"is_active"`
	AddTime          string  `json:"add_time"`
	RemoveTime       *string `json:"remove_time,omitempty"`
	Type             string  `json:"type,omitempty"`
	HTTPAuthUser     *string `json:"http_auth_user,omitempty"`
	HTTPAuthPassword *string `json:"http_auth_password,omitempty"`
	AdditionalData   *string `json:"additional_data,omitempty"`
	RemoveReason     *string `json:"remove_reason,omitempty"`
	LastDeliveryTime *string `json:"last_delivery_time,omitempty"`
	LastHTTPStatus   *string `json:"last_http_status,omitempty"`
	AdminID          int     `json:"admin_id,omitempty"`
}

type WebhooksResponse struct {
	Success bool      `json:"success"`
	Data    []Webhook `json:"data"`
}

type WebhookResponse struct {
	Success bool    `json:"success"`
	Data    Webhook `json:"data"`
}

type CreateWebhookReq struct {
	SubscriptionURL  string  `json:"subscription_url"`
	EventAction      string  `json:"event_action"`
	EventObject      string  `json:"event_object"`
	UserID           *int    `json:"user_id,omitempty"`
	HTTPAuthUser     *string `json:"http_auth_user,omitempty"`
	HTTPAuthPassword *string `json:"http_auth_password,omitempty"`
	Version          *string `json:"version,omitempty"`
}
