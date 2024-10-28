package models

type Activity struct {
	ID                         int      `json:"id"`
	CompanyID                  int      `json:"company_id"`
	UserID                     int      `json:"user_id"`
	Done                       bool     `json:"done"`
	Type                       string   `json:"type"`
	ReferenceType              string   `json:"reference_type"`
	ReferenceID                int      `json:"reference_id"`
	ConferenceMeetingClient    string   `json:"conference_meeting_client"`
	ConferenceMeetingURL       string   `json:"conference_meeting_url"`
	ConferenceMeetingID        string   `json:"conference_meeting_id"`
	DueDate                    string   `json:"due_date"`
	DueTime                    string   `json:"due_time"`
	Duration                   string   `json:"duration"`
	BusyFlag                   bool     `json:"busy_flag"`
	AddTime                    string   `json:"add_time"`
	MarkedAsDoneTime           string   `json:"marked_as_done_time"`
	LastNotificationTime       string   `json:"last_notification_time"`
	LastNotificationUserID     int      `json:"last_notification_user_id"`
	NotificationLanguageID     int      `json:"notification_language_id"`
	Subject                    string   `json:"subject"`
	PublicDescription          string   `json:"public_description"`
	CalendarSyncIncludeContext string   `json:"calendar_sync_include_context"`
	Location                   string   `json:"location"`
	OrgID                      int      `json:"org_id"`
	PersonID                   int      `json:"person_id"`
	DealID                     int      `json:"deal_id"`
	LeadID                     string   `json:"lead_id"`
	ProjectID                  int      `json:"project_id"`
	ActiveFlag                 bool     `json:"active_flag"`
	UpdateTime                 string   `json:"update_time"`
	UpdateUserID               int      `json:"update_user_id"`
	GcalEventID                string   `json:"gcal_event_id"`
	GoogleCalendarID           string   `json:"google_calendar_id"`
	GoogleCalendarEtag         string   `json:"google_calendar_etag"`
	SourceTimezone             string   `json:"source_timezone"`
	RecRule                    string   `json:"rec_rule"`
	RecRuleExtension           string   `json:"rec_rule_extension"`
	RecMasterActivityID        int      `json:"rec_master_activity_id"`
	Series                     []string `json:"series"`
	Note                       string   `json:"note"`
	CreatedByUserID            int      `json:"created_by_user_id"`
	LocationSubpremise         string   `json:"location_subpremise"`
	LocationStreetNumber       string   `json:"location_street_number"`
	LocationRoute              string   `json:"location_route"`
	LocationSublocality        string   `json:"location_sublocality"`
	LocationLocality           string   `json:"location_locality"`
	PersonDropboxBCC           string   `json:"person_dropbox_bcc"`
	DealDropboxBCC             string   `json:"deal_dropbox_bcc"`
	AssignedToUserID           int      `json:"assigned_to_user_id"`
	CleanName                  string   `json:"clean_name"`
	URL                        string   `json:"url"`
	Call                       int      `json:"call"`
	Meeting                    int      `json:"meeting"`
	Name                       string   `json:"name"`
	ActivityCount              int      `json:"activity_count"`
	Share                      int      `json:"share"`
}

type ActivityResponse struct {
	Success bool      `json:"success"`
	Data    Activity  `json:"data"`
}

type ActivitiesResponse struct {
	Success bool      `json:"success"`
	Data    []Activity `json:"data"`
}

type ActivityRequest struct {
	DueDate           string `json:"due_date"`
	DueTime           string `json:"due_time"`
	Duration          string `json:"duration"`
	DealID            int    `json:"deal_id,omitempty"`
	LeadID            string `json:"lead_id,omitempty"`
	PersonID          int    `json:"person_id,omitempty"`
	OrgID             int    `json:"org_id,omitempty"`
	Location          string `json:"location,omitempty"`
	PublicDescription string `json:"public_description,omitempty"`
	Note              string `json:"note,omitempty"`
	Subject           string `json:"subject,omitempty"`
	Type              string `json:"type,omitempty"`
	UserID            int    `json:"user_id,omitempty"`
	BusyFlag          bool   `json:"busy_flag,omitempty"`
	Done              int    `json:"done,omitempty"`
}

type ActivityUpdateRequest struct {
	DueDate           string `json:"due_date,omitempty"`
	DueTime           string `json:"due_time,omitempty"`
	Duration          string `json:"duration,omitempty"`
	Location          string `json:"location,omitempty"`
	PublicDescription string `json:"public_description,omitempty"`
	Note              string `json:"note,omitempty"`
	Subject           string `json:"subject,omitempty"`
	Type              string `json:"type,omitempty"`
	UserID            int    `json:"user_id,omitempty"`
	Done              int    `json:"done,omitempty"`
}

type ActivitiesBetaResponse struct {
	Success       bool           `json:"success"`
	Data          []Activity     `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}
