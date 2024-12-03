package models

type MailMessageResponse struct {
	Success        bool          `json:"success"`
	StatusCode     int           `json:"statusCode"`
	StatusText     string        `json:"statusText"`
	Service        string        `json:"service"`
	MailMessage    MailMessage   `json:"mail_message"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type MailMessage struct {
	ID                       int       `json:"id"`
	EmailAddress             string    `json:"email_address"`
	Name                     string    `json:"name"`
	LinkedPersonID           int       `json:"linked_person_id"`
	LinkedPersonName         string    `json:"linked_person_name"`
	MailMessagePartyID       int       `json:"mail_message_party_id"`
	BodyURL                  string    `json:"body_url"`
	AccountID                string    `json:"account_id"`
	UserID                   int       `json:"user_id"`
	MailThreadID             int       `json:"mail_thread_id"`
	Subject                  string    `json:"subject"`
	Snippet                  string    `json:"snippet"`
	MailTrackingStatus       string    `json:"mail_tracking_status"`
	MailLinkTrackingEnabled  int       `json:"mail_link_tracking_enabled_flag"`
	ReadFlag                 int       `json:"read_flag"`
	Draft                    string    `json:"draft"`
	DraftFlag                int       `json:"draft_flag"`
	SyncedFlag               int       `json:"synced_flag"`
	DeletedFlag              int       `json:"deleted_flag"`
	HasBodyFlag              int       `json:"has_body_flag"`
	SentFlag                 int       `json:"sent_flag"`
	SentFromPipedriveFlag    int       `json:"sent_from_pipedrive_flag"`
	SmartBCCFlag             int       `json:"smart_bcc_flag"`
	MessageTime              string    `json:"message_time"`
	AddTime                  string    `json:"add_time"`
	UpdateTime               string    `json:"update_time"`
	HasAttachmentsFlag       int       `json:"has_attachments_flag"`
	HasInlineAttachmentsFlag int       `json:"has_inline_attachments_flag"`
	HasRealAttachmentsFlag   int       `json:"has_real_attachments_flag"`
}

type MailThreadResponse struct {
	Success        bool         `json:"success"`
	StatusCode     int          `json:"statusCode"`
	StatusText     string       `json:"statusText"`
	Service        string       `json:"service"`
	MailThread     MailThread   `json:"mail_thread"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type MailThread struct {
	ID                       int       `json:"id"`
	EmailAddress             string    `json:"email_address"`
	Name                     string    `json:"name"`
	MessageTime              int64     `json:"message_time"`
	LinkedPersonID           int       `json:"linked_person_id"`
	LinkedPersonName         string    `json:"linked_person_name"`
	Subject                  string    `json:"subject"`
	Snippet                  string    `json:"snippet"`
	ReadFlag                 int       `json:"read_flag"`
	MailTrackingStatus       string    `json:"mail_tracking_status"`
	HasAttachmentsFlag       int       `json:"has_attachments_flag"`
	ArchivedFlag             int       `json:"archived_flag"`
	DeletedFlag              int       `json:"deleted_flag"`
	SyncedFlag               int       `json:"synced_flag"`
	FirstMessageTimestamp    string    `json:"first_message_timestamp"`
	LastMessageTimestamp     string    `json:"last_message_timestamp"`
	AddTime                  string    `json:"add_time"`
	UpdateTime               string    `json:"update_time"`
	DealID                   int       `json:"deal_id"`
	DealStatus               string    `json:"deal_status"`
}
