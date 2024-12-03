package models

// Note represents a single note.
type Note struct {
	ID                       int      `json:"id"`
	ActiveFlag               bool     `json:"active_flag"`
	AddTime                  string   `json:"add_time"`
	Content                  string   `json:"content"`
	Title                    string   `json:"title"`
	LeadID                   string   `json:"lead_id,omitempty"`
	DealID                   int      `json:"deal_id,omitempty"`
	OrgID                    int      `json:"org_id,omitempty"`
	PersonID                 int      `json:"person_id,omitempty"`
	UserID                   int      `json:"user_id"`
	UpdateTime               string   `json:"update_time"`
	PinnedToLeadFlag         bool     `json:"pinned_to_lead_flag,omitempty"`
	PinnedToDealFlag         bool     `json:"pinned_to_deal_flag,omitempty"`
	PinnedToOrganizationFlag bool     `json:"pinned_to_organization_flag,omitempty"`
	PinnedToPersonFlag       bool     `json:"pinned_to_person_flag,omitempty"`
}

// NotesResponse represents a response with a list of notes.
type NotesResponse struct {
	Success bool   `json:"success"`
	Data    []Note `json:"data"`
}

// SingleNoteResponse represents a response with a single note.
type SingleNoteResponse struct {
	Success bool `json:"success"`
	Data    Note `json:"data"`
}

// Comment represents a comment associated with a note.
type Comment struct {
	UUID        string `json:"uuid"`
	ActiveFlag  bool   `json:"active_flag"`
	AddTime     string `json:"add_time"`
	Content     string `json:"content"`
	UpdateTime  string `json:"update_time"`
	UserID      int    `json:"user_id"`
	ObjectID    string `json:"object_id"`
	ObjectType  string `json:"object_type"`
	UpdaterID   int    `json:"updater_id"`
	CompanyID   int    `json:"company_id"`
}

// CommentsResponse represents a response with a list of comments for a note.
type CommentsResponse struct {
	Success bool      `json:"success"`
	Data    []Comment `json:"data"`
}

// SingleCommentResponse represents a response with a single comment.
type SingleCommentResponse struct {
	Success bool    `json:"success"`
	Data    Comment `json:"data"`
}

// AddNoteRequest represents a request body for adding a note.
type AddNoteRequest struct {
	Content                  string `json:"content"`
	LeadID                   string `json:"lead_id,omitempty"`
	DealID                   int    `json:"deal_id,omitempty"`
	PersonID                 int    `json:"person_id,omitempty"`
	OrgID                    int    `json:"org_id,omitempty"`
	UserID                   int    `json:"user_id,omitempty"`
	AddTime                  string `json:"add_time,omitempty"`
	PinnedToLeadFlag         int    `json:"pinned_to_lead_flag,omitempty"`
	PinnedToDealFlag         int    `json:"pinned_to_deal_flag,omitempty"`
	PinnedToOrganizationFlag int    `json:"pinned_to_organization_flag,omitempty"`
	PinnedToPersonFlag       int    `json:"pinned_to_person_flag,omitempty"`
}

// AddCommentRequest represents a request body for adding a comment to a note.
type AddCommentRequest struct {
	Content string `json:"content"`
}
