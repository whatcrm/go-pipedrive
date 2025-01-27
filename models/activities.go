package models

type Activity struct {
	ID                int            `json:"id"`
	Subject           string         `json:"subject"`
	OwnerID           int            `json:"owner_id"`
	Type              string         `json:"type"`
	IsDeleted         bool           `json:"is_deleted"`
	Done              bool           `json:"done"`
	DueDate           string         `json:"due_date"`
	DueTime           string         `json:"due_time"`
	Duration          string         `json:"duration"`
	Busy              bool           `json:"busy"`
	AddTime           string         `json:"add_time"`
	UpdateTime        string         `json:"update_time"`
	MarkedAsDoneTime  string         `json:"marked_as_done_time"`
	PublicDescription string         `json:"public_description"`
	Location          Location       `json:"location"`
	OrgID             int            `json:"org_id"`
	PersonID          int            `json:"person_id"`
	DealID            int            `json:"deal_id"`
	LeadID            string         `json:"lead_id"`
	ProjectID         int            `json:"project_id"`
	Private           bool           `json:"private"`
	Priority          int            `json:"priority"`
	Note              string         `json:"note"`
	CreatedUserID     int            `json:"created_user_id"`
	Attendees         []Attendees    `json:"attendees,omitempty"`
	Participants      []Participants `json:"participants,omitempty"`
}

type Attendees struct {
	EmailAddress string `json:"email_address"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	IsOrganizer  bool   `json:"is_organizer"`
	PersonID     int    `json:"person_id"`
	UserID       int    `json:"user_id"`
}

type Participants struct {
	PersonID    int  `json:"person_id"`
	PrimaryFlag bool `json:"primary_flag"`
}

type Location struct {
	Value            string `json:"value"`
	StreetNumber     string `json:"street_number"`
	Route            string `json:"route"`
	Sublocality      string `json:"sublocality"`
	Locality         string `json:"locality"`
	AdminAreaLevel1  string `json:"admin_area_level_1"`
	AdminAreaLevel2  string `json:"admin_area_level_2"`
	Country          string `json:"country"`
	PostalCode       string `json:"postal_code"`
	FormattedAddress string `json:"formatted_address"`
}

type ActivityResponse struct {
	Success bool     `json:"success"`
	Data    Activity `json:"data"`
}

type ActivitiesResponse struct {
	Success bool       `json:"success"`
	Data    []Activity `json:"data"`
}

type ActivityRequest struct {
	DueDate           string   `json:"due_date"`
	DueTime           string   `json:"due_time"`
	Duration          string   `json:"duration"`
	DealID            int      `json:"deal_id,omitempty"`
	LeadID            string   `json:"lead_id,omitempty"`
	PersonID          int      `json:"person_id,omitempty"`
	OrgID             int      `json:"org_id,omitempty"`
	Location          Location `json:"location,omitempty"`
	PublicDescription string   `json:"public_description,omitempty"`
	Note              string   `json:"note,omitempty"`
	Subject           string   `json:"subject,omitempty"`
	Type              string   `json:"type,omitempty"`
	OwnerID           int      `json:"owner_id,omitempty"`
	Busy              bool     `json:"busy,omitempty"`
	Done              bool     `json:"done,omitempty"`
}

type ActivityUpdateRequest struct {
	DueDate           string   `json:"due_date,omitempty"`
	DueTime           string   `json:"due_time,omitempty"`
	Duration          string   `json:"duration,omitempty"`
	Location          Location `json:"location,omitempty"`
	PublicDescription string   `json:"public_description,omitempty"`
	Note              string   `json:"note,omitempty"`
	Subject           string   `json:"subject,omitempty"`
	Type              string   `json:"type,omitempty"`
	OwnerID           int      `json:"owner_id,omitempty"`
	Done              bool     `json:"done,omitempty"`
}

type ActivitiesBetaResponse struct {
	Success        bool       `json:"success"`
	Data           []Activity `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}
