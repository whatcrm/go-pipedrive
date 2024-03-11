package models

import "time"

type Webhook struct {
    ID                int        `json:"id"`
    CompanyID         int        `json:"company_id"`
    OwnerID           int        `json:"owner_id"`
    UserID            int        `json:"user_id"`
    EventAction       string     `json:"event_action"`
    EventObject       string     `json:"event_object"`
    SubscriptionURL   string     `json:"subscription_url"`
    Version           string     `json:"version"`
    IsActive          int        `json:"is_active"`
    AddTime           time.Time  `json:"add_time"`
    RemoveTime        *time.Time `json:"remove_time"`
    Type              string     `json:"type"`
    HTTPAuthUser      *string    `json:"http_auth_user,omitempty"`
    HTTPAuthPassword  *string    `json:"http_auth_password,omitempty"`
    AdditionalData    string     `json:"additional_data"`
    RemoveReason      *string    `json:"remove_reason,omitempty"`
    LastDeliveryTime  *time.Time `json:"last_delivery_time,omitempty"`
    LastHTTPStatus    *int       `json:"last_http_status,omitempty"`
    AdminID           int        `json:"admin_id"`
}