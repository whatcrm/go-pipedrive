package models

type CallLogsResponse struct {
	Success bool      `json:"success"`
	Data    []CallLog `json:"data"`
}

type CallLogResponse struct {
	Success bool    `json:"success"`
	Data    CallLog `json:"data"`
}

type CallLog struct {
	ID              string  `json:"id"`
	ActivityID      int     `json:"activity_id"`
	PersonID        *int    `json:"person_id"`
	OrgID           *int    `json:"org_id"`
	DealID          *int    `json:"deal_id"`
	LeadID          *string `json:"lead_id"`
	Subject         string  `json:"subject"`
	Duration        string  `json:"duration"`
	Outcome         string  `json:"outcome"`
	FromPhoneNumber string  `json:"from_phone_number"`
	ToPhoneNumber   string  `json:"to_phone_number"`
	HasRecording    bool    `json:"has_recording"`
	StartTime       string  `json:"start_time"`
	EndTime         string  `json:"end_time"`
	UserID          int     `json:"user_id"`
	CompanyID       int     `json:"company_id"`
	Note            string  `json:"note"`
}

type CallLogReq struct {
	ActivityID      int     `json:"activity_id,omitempty"`
	PersonID        *int    `json:"person_id,omitempty"`
	OrgID           *int    `json:"org_id,omitempty"`
	DealID          *int    `json:"deal_id,omitempty"`
	LeadID          *string `json:"lead_id,omitempty"`
	Subject         string  `json:"subject,omitempty"`
	Duration        string  `json:"duration,omitempty"`
	Outcome         string  `json:"outcome"` //required values: connected, no_answer, left_message, left_voicemail, wrong_number, busy
	FromPhoneNumber string  `json:"from_phone_number,omitempty"`
	ToPhoneNumber   string  `json:"to_phone_number"` //required 
	StartTime       string  `json:"start_time"`//required 
	EndTime         string  `json:"end_time"`//required 
	UserID          int     `json:"user_id,omitempty"`
	Note            string  `json:"note,omitempty"`
}
