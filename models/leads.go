package models

type Lead struct {
	ID                string  `json:"id"`
	Title             string  `json:"title"`
	OwnerID           int     `json:"owner_id"`
	CreatorID         int     `json:"creator_id"`
	PersonID          int     `json:"person_id,omitempty"`
	OrganizationID    int     `json:"organization_id,omitempty"`
	SourceName        string  `json:"source_name"`
	Origin            string  `json:"origin"`
	OriginID          string  `json:"origin_id,omitempty"`
	Channel           int     `json:"channel"`
	ChannelID         string  `json:"channel_id,omitempty"`
	IsArchived        bool    `json:"is_archived"`
	WasSeen           bool    `json:"was_seen"`
	Amount            float64 `json:"amount,omitempty"`
	Currency          string  `json:"currency,omitempty"`
	ExpectedCloseDate string  `json:"expected_close_date,omitempty"`
	NextActivityID    int     `json:"next_activity_id,omitempty"`
	AddTime           string  `json:"add_time"`
	UpdateTime        string  `json:"update_time"`
	VisibleTo         string  `json:"visible_to"`
	CCEmail           string  `json:"cc_email"`
}

type LeadResponse struct {
	Success bool `json:"success"`
	Data    Lead `json:"data"`
}

type LeadsResponse struct {
	Success bool   `json:"success"`
	Data    []Lead `json:"data"`
}

type PermittedUsersResponse struct {
	Success bool  `json:"success"`
	Data    []int `json:"data"`
}

type LeadSearchResponse struct {
	Success bool   `json:"success"`
	Data    []Lead `json:"data"`
}

type LeadRequest struct {
	Title          string   `json:"title"`
	OwnerID        int      `json:"owner_id,omitempty"`
	LabelIDs       []string `json:"label_ids,omitempty"`
	PersonID       int      `json:"person_id,omitempty"`
	OrganizationID int      `json:"organization_id,omitempty"`
	Value          struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"value,omitempty"`
	ExpectedCloseDate string `json:"expected_close_date,omitempty"`
	VisibleTo         string `json:"visible_to,omitempty"`
	WasSeen           bool   `json:"was_seen,omitempty"`
	OriginID          string `json:"origin_id,omitempty"`
	Channel           int    `json:"channel,omitempty"`
	ChannelID         string `json:"channel_id,omitempty"`
}

type LeadUpdateRequest struct {
	Title          string `json:"title,omitempty"`
	OwnerID        int    `json:"owner_id,omitempty"`
	LabelIDs       []int  `json:"label_ids,omitempty"`
	PersonID       int    `json:"person_id,omitempty"`
	OrganizationID int    `json:"organization_id,omitempty"`
	IsArchived     bool   `json:"is_archived,omitempty"`
	Value          struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"value,omitempty"`
	ExpectedCloseDate string `json:"expected_close_date,omitempty"`
	VisibleTo         string `json:"visible_to,omitempty"`
	WasSeen           bool   `json:"was_seen,omitempty"`
	Channel           int    `json:"channel,omitempty"`
	ChannelID         string `json:"channel_id,omitempty"`
}
