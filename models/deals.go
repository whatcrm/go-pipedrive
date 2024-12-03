package models

type Deal struct {
	ID                int     `json:"id"`
	Title             string  `json:"title"`
	Value             float64 `json:"value"`
	Currency          string  `json:"currency"`
	UserID            int     `json:"user_id"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	Status            string  `json:"status"`
	AddTime           string  `json:"add_time"`
	UpdateTime        string  `json:"update_time"`
}

type DealRequest struct {
	Title             string  `json:"title"`
	Value             float64 `json:"value,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	UserID            int     `json:"user_id,omitempty"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	PipelineID        int     `json:"pipeline_id,omitempty"`
	StageID           int     `json:"stage_id,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	Status            string  `json:"status,omitempty"`
}

type DealUpdateRequest struct {
	Title             string  `json:"title,omitempty"`
	Value             float64 `json:"value,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	UserID            int     `json:"user_id,omitempty"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	PipelineID        int     `json:"pipeline_id,omitempty"`
	StageID           int     `json:"stage_id,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	Status            string  `json:"status,omitempty"`
}

type DealResponse struct {
	Success bool `json:"success"`
	Data    Deal `json:"data"`
}

type DealsResponse struct {
	Success bool   `json:"success"`
	Data    []Deal `json:"data"`
}

type Participant struct {
	ID              int             `json:"id"`
	AddTime         string          `json:"add_time"`
	ActiveFlag      bool            `json:"active_flag"`
	RelatedItemData RelatedItemData `json:"related_item_data"`
	Person          Person          `json:"person"`
	AddedByUserID   User            `json:"added_by_user_id"`
	RelatedItemType string          `json:"related_item_type"`
	RelatedItemID   int             `json:"related_item_id"`
}
type RelatedItemData struct {
	DealID int    `json:"deal_id"`
	Title  string `json:"title"`
}

// DealParticipantsResponse represents the response from the API when fetching deal participants.
type ParticipantsResponse struct {
	Success        bool          `json:"success"`
	Data           []Participant `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}
