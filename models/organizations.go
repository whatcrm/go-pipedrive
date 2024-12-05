package models

// Organization represents an organization entity.
type Organization struct {
	ID           int                    `json:"id"`
	Name         string                 `json:"name"`
	PeopleCount  int                    `json:"people_count"`
	OwnerID      int                    `json:"owner_id"`
	Address      string                 `json:"address"`
	ActiveFlag   bool                   `json:"active_flag"`
	CCEmail      string                 `json:"cc_email"`
	VisibleTo    int                    `json:"visible_to"`
	AddTime      string                 `json:"add_time"`
	UpdateTime   string                 `json:"update_time"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// OrganizationResponse represents the response for an organization.
type OrganizationResponse struct {
	Success bool         `json:"success"`
	Data    Organization `json:"data"`
}

// OrganizationsResponse represents the response for multiple organizations.
type OrganizationsResponse struct {
	Success               bool           `json:"success"`
	Data                  []Organization `json:"data"`
	MoreItemsInCollection bool           `json:"more_items_in_collection"`
	NextStart             int            `json:"next_start,omitempty"`
}

// FollowerResponse represents the response for follower actions.
type FollowerResponse struct {
	Success bool `json:"success"`
	ID      int  `json:"id"`
}
