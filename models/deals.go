package models

// Deal represents a single deal in the Pipedrive system.
type Deal struct {
	ID                     int          `json:"id"`
	CreatorUserID          User         `json:"creator_user_id"`
	UserID                 User         `json:"user_id"`
	PersonID               DealPerson   `json:"person_id"`
	OrgID                  Organization `json:"org_id"`
	StageID                int          `json:"stage_id"`
	Title                  string       `json:"title"`
	Value                  float64      `json:"value"`
	Currency               string       `json:"currency"`
	AddTime                string       `json:"add_time"`
	UpdateTime             string       `json:"update_time"`
	StageChangeTime        string       `json:"stage_change_time"`
	Active                 bool         `json:"active"`
	Deleted                bool         `json:"deleted"`
	Status                 string       `json:"status"`
	Probability            float64      `json:"probability"`
	NextActivityDate       string       `json:"next_activity_date"`
	NextActivityTime       string       `json:"next_activity_time"`
	NextActivityID         int          `json:"next_activity_id"`
	LastActivityID         int          `json:"last_activity_id"`
	LastActivityDate       string       `json:"last_activity_date"`
	LostReason             string       `json:"lost_reason"`
	VisibleTo              string       `json:"visible_to"`
	CloseTime              string       `json:"close_time"`
	PipelineID             int          `json:"pipeline_id"`
	WonTime                string       `json:"won_time"`
	FirstWonTime           string       `json:"first_won_time"`
	LostTime               string       `json:"lost_time"`
	ProductsCount          int          `json:"products_count"`
	FilesCount             int          `json:"files_count"`
	NotesCount             int          `json:"notes_count"`
	FollowersCount         int          `json:"followers_count"`
	EmailMessagesCount     int          `json:"email_messages_count"`
	ActivitiesCount        int          `json:"activities_count"`
	DoneActivitiesCount    int          `json:"done_activities_count"`
	UndoneActivitiesCount  int          `json:"undone_activities_count"`
	ParticipantsCount      int          `json:"participants_count"`
	ExpectedCloseDate      string       `json:"expected_close_date"`
	LastIncomingMailTime   string       `json:"last_incoming_mail_time"`
	LastOutgoingMailTime   string       `json:"last_outgoing_mail_time"`
	Label                  string       `json:"label"`
	LocalWonDate           string       `json:"local_won_date"`
	LocalLostDate          string       `json:"local_lost_date"`
	LocalCloseDate         string       `json:"local_close_date"`
	Origin                 string       `json:"origin"`
	OriginID               string       `json:"origin_id"`
	Channel                string       `json:"channel"`
	ChannelID              string       `json:"channel_id"`
	StageOrderNr           int          `json:"stage_order_nr"`
	PersonName             string       `json:"person_name"`
	OrgName                string       `json:"org_name"`
	NextActivitySubject    string       `json:"next_activity_subject"`
	NextActivityType       string       `json:"next_activity_type"`
	NextActivityDuration   string       `json:"next_activity_duration"`
	NextActivityNote       string       `json:"next_activity_note"`
	FormattedValue         string       `json:"formatted_value"`
	WeightedValue          float64      `json:"weighted_value"`
	FormattedWeightedValue string       `json:"formatted_weighted_value"`
	WeightedValueCurrency  string       `json:"weighted_value_currency"`
	RottenTime             string       `json:"rotten_time"`
	ACVCurrency            string       `json:"acv_currency"`
	MRRCurrency            string       `json:"mrr_currency"`
	ARRCurrency            string       `json:"arr_currency"`
	OwnerName              string       `json:"owner_name"`
	CCEmail                string       `json:"cc_email"`
	OrgHidden              bool         `json:"org_hidden"`
	PersonHidden           bool         `json:"person_hidden"`
}

// DealsResponse represents the response from the API when fetching multiple deals.
type DealsResponse struct {
	Success        bool   `json:"success"`
	Data           []Deal `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}

// DealResponse represents the response from the API when fetching a single deal.
type DealResponse struct {
	Success bool `json:"success"`
	Data    Deal `json:"data"`
}

// Person represents a person in the Pipedrive system.
type PersonsRequest struct {
	Name         string                 `json:"name"`
	OwnerID      int                    `json:"owner_id,omitempty"`
	OrgID        int                    `json:"org_id,omitempty"`
	Email        []string               `json:"email,omitempty"`
	Phone        []string               `json:"phone,omitempty"`
	VisibleTo    string                 `json:"visible_to,omitempty"`
	AddTime      string                 `json:"add_time,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
}

// Person represents a person in the Pipedrive system.
type Person struct {
	ID         int          `json:"id"`
	ActiveFlag bool         `json:"active_flag"`
	Name       string       `json:"name"`
	Email      []Email      `json:"email"`
	Phone      []Phone      `json:"phone"`
	OwnerID    Owner        `json:"owner_id"`
	OrgID      Organization `json:"org_id"`
}

type DealPerson struct {
	ID         int          `json:"id"`
	ActiveFlag bool         `json:"active_flag"`
	Name       string       `json:"name"`
	Email      []Email      `json:"email"`
	Phone      []Phone      `json:"phone"`
	OwnerID    int          `json:"owner_id"`
	OrgID      Organization `json:"org_id"`
}

type Owner struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	HasPic     int    `json:"has_pic"`
	PicHash    string `json:"pic_hash"`
	ActiveFlag bool   `json:"active_flag"`
	Value      int    `json:"value"`
}

type Organization struct {
	Name        string `json:"name"`
	PeopleCount int    `json:"people_count"`
	OwnerID     int    `json:"owner_id"`
	Address     string `json:"address"`
	ActiveFlag  bool   `json:"active_flag"`
	CCEmail     string `json:"cc_email"`
	LabelIDs    []int  `json:"label_ids"`
	OwnerName   string `json:"owner_name"`
	Value       int    `json:"value"`
}

// Email represents an email in the Pipedrive system.
type Email struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// Phone represents a phone number in the Pipedrive system.
type Phone struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}
type DealRequest struct {
	Title             string `json:"title"`
	Value             int    `json:"value,omitempty"`
	Currency          string `json:"currency,omitempty"`
	UserID            int    `json:"user_id,omitempty"`
	PersonID          int    `json:"person_id,omitempty"`
	OrgID             int    `json:"org_id,omitempty"`
	PipelineID        int    `json:"pipeline_id,omitempty"`
	StageID           int    `json:"stage_id,omitempty"`
	Status            string `json:"status,omitempty"`
	ExpectedCloseDate string `json:"expected_close_date,omitempty"`
}

type DealUpdateRequest struct {
	Title             string `json:"title,omitempty"`
	Value             int    `json:"value,omitempty"`
	Currency          string `json:"currency,omitempty"`
	UserID            int    `json:"user_id,omitempty"`
	PersonID          int    `json:"person_id,omitempty"`
	OrgID             int    `json:"org_id,omitempty"`
	PipelineID        int    `json:"pipeline_id,omitempty"`
	StageID           int    `json:"stage_id,omitempty"`
	Status            string `json:"status,omitempty"`
	ExpectedCloseDate string `json:"expected_close_date,omitempty"`
}

type DealFollower struct {
	ID     int `json:"id"`
	UserID int `json:"user_id"`
	DealID int `json:"deal_id"`
}

type DealFollowerRequest struct {
	UserID int `json:"user_id"`
}

type DealFollowerResponse struct {
	Success bool         `json:"success"`
	Data    DealFollower `json:"data"`
}

type DealParticipantRequest struct {
	PersonID int `json:"person_id"`
}

// DealParticipant represents a participant in a deal.
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

type DealProduct struct {
	ID           int     `json:"id"`
	DealID       int     `json:"deal_id"`
	ProductID    int     `json:"product_id"`
	ItemPrice    float64 `json:"item_price"`
	Quantity     int     `json:"quantity"`
	Discount     float64 `json:"discount"`
	Tax          float64 `json:"tax"`
	Comments     string  `json:"comments"`
	IsEnabled    bool    `json:"is_enabled"`
	TaxMethod    string  `json:"tax_method"`
	DiscountType string  `json:"discount_type"`
}

type DealProductRequest struct {
	ProductID    int     `json:"product_id"`
	ItemPrice    float64 `json:"item_price"`
	Quantity     int     `json:"quantity"`
	Discount     float64 `json:"discount,omitempty"`
	Tax          float64 `json:"tax,omitempty"`
	Comments     string  `json:"comments,omitempty"`
	IsEnabled    bool    `json:"is_enabled,omitempty"`
	TaxMethod    string  `json:"tax_method,omitempty"`
	DiscountType string  `json:"discount_type,omitempty"`
}

type DealProductUpdateRequest struct {
	ProductID    int     `json:"product_id,omitempty"`
	ItemPrice    float64 `json:"item_price,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Discount     float64 `json:"discount,omitempty"`
	Tax          float64 `json:"tax,omitempty"`
	Comments     string  `json:"comments,omitempty"`
	IsEnabled    bool    `json:"is_enabled,omitempty"`
	TaxMethod    string  `json:"tax_method,omitempty"`
	DiscountType string  `json:"discount_type,omitempty"`
}

type DealProductResponse struct {
	Success bool        `json:"success"`
	Data    DealProduct `json:"data"`
}

type DealDiscount struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	DealID      int     `json:"deal_id"`
}

type DealDiscountRequest struct {
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
}

type DealDiscountUpdateRequest struct {
	Description string  `json:"description,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	Type        string  `json:"type,omitempty"`
}

type DealDiscountResponse struct {
	Success bool         `json:"success"`
	Data    DealDiscount `json:"data"`
}

type MergeDealRequest struct {
	MergeWithID int `json:"merge_with_id"`
}

type DealsSummaryResponse struct {
	Success                                      bool               `json:"success"`
	Data                                         []DealsSummaryData `json:"data"`
	TotalCount                                   int                `json:"total_count"`
	TotalCurrencyConvertedValue                  float64            `json:"total_currency_converted_value"`
	TotalWeightedCurrencyConvertedValue          float64            `json:"total_weighted_currency_converted_value"`
	TotalCurrencyConvertedValueFormatted         string             `json:"total_currency_converted_value_formatted"`
	TotalWeightedCurrencyConvertedValueFormatted string             `json:"total_weighted_currency_converted_value_formatted"`
}

type DealsSummaryData struct {
	Value                   float64 `json:"value"`
	Count                   int     `json:"count"`
	ValueConverted          float64 `json:"value_converted"`
	ValueFormatted          string  `json:"value_formatted"`
	ValueConvertedFormatted string  `json:"value_converted_formatted"`
}

type DealsTimelineResponse struct {
	Success bool                `json:"success"`
	Data    []DealsTimelineData `json:"data"`
}

type DealsTimelineData struct {
	PeriodStart string `json:"period_start"`
	PeriodEnd   string `json:"period_end"`
	Deals       []Deal `json:"deals"`
}

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

type ActivitiesResponse struct {
	Success bool       `json:"success"`
	Data    []Activity `json:"data"`
}

type DealFieldUpdate struct {
	FieldKey              string `json:"field_key"`
	OldValue              string `json:"old_value"`
	NewValue              string `json:"new_value"`
	ActorUserID           int    `json:"actor_user_id"`
	Time                  string `json:"time"`
	ChangeSource          string `json:"change_source"`
	ChangeSourceUserAgent string `json:"change_source_user_agent"`
	IsBulkUpdateFlag      bool   `json:"is_bulk_update_flag"`
}

type DealFieldUpdatesResponse struct {
	Success bool              `json:"success"`
	Data    []DealFieldUpdate `json:"data"`
}

type DealUpdate struct {
	Object                string `json:"object"`
	Timestamp             string `json:"timestamp"`
	ID                    int    `json:"id"`
	ItemID                int    `json:"item_id"`
	UserID                int    `json:"user_id"`
	FieldKey              string `json:"field_key"`
	OldValue              string `json:"old_value"`
	NewValue              string `json:"new_value"`
	IsBulkUpdateFlag      bool   `json:"is_bulk_update_flag"`
	LogTime               string `json:"log_time"`
	ChangeSource          string `json:"change_source"`
	ChangeSourceUserAgent string `json:"change_source_user_agent"`
	AdditionalData        string `json:"additional_data"`
}

type DealUpdatesResponse struct {
	Success bool         `json:"success"`
	Data    []DealUpdate `json:"data"`
}

type Follower struct {
	UserID  int    `json:"user_id"`
	ID      int    `json:"id"`
	DealID  int    `json:"deal_id"`
	AddTime string `json:"add_time"`
}

type FollowersResponse struct {
	Success bool       `json:"success"`
	Data    []Follower `json:"data"`
}

type MailMessage struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	DealID         int    `json:"deal_id"`
	PersonID       int    `json:"person_id"`
	OrgID          int    `json:"org_id"`
	ProductID      int    `json:"product_id"`
	ActivityID     int    `json:"activity_id"`
	LeadID         string `json:"lead_id"`
	LogID          int    `json:"log_id"`
	AddTime        string `json:"add_time"`
	UpdateTime     string `json:"update_time"`
	FileName       string `json:"file_name"`
	FileType       string `json:"file_type"`
	FileSize       int    `json:"file_size"`
	ActiveFlag     bool   `json:"active_flag"`
	InlineFlag     bool   `json:"inline_flag"`
	RemoteLocation string `json:"remote_location"`
	RemoteID       string `json:"remote_id"`
	CID            string `json:"cid"`
	S3Bucket       string `json:"s3_bucket"`
	MailMessageID  string `json:"mail_message_id"`
	MailTemplateID string `json:"mail_template_id"`
	DealName       string `json:"deal_name"`
	PersonName     string `json:"person_name"`
	OrgName        string `json:"org_name"`
	ProductName    string `json:"product_name"`
	LeadName       string `json:"lead_name"`
	URL            string `json:"url"`
}

type MailMessagesResponse struct {
	Success bool          `json:"success"`
	Data    []MailMessage `json:"data"`
}

type PersonsResponse struct {
	Success        bool     `json:"success"`
	Data           []Person `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}

type PersonResponse struct {
	Success        bool   `json:"success"`
	Data           Person `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}

type Product struct {
	ID                     int     `json:"id"`
	DealID                 int     `json:"deal_id"`
	ProductID              int     `json:"product_id"`
	ProductName            string  `json:"product_name"`
	Comments               string  `json:"comments"`
	ItemPrice              float64 `json:"item_price"`
	Quantity               int     `json:"quantity"`
	Discount               float64 `json:"discount"`
	DiscountType           string  `json:"discount_type"`
	Tax                    float64 `json:"tax"`
	TaxMethod              string  `json:"tax_method"`
	IsEnabled              bool    `json:"is_enabled"`
	AddTime                string  `json:"add_time"`
	UpdateTime             string  `json:"update_time"`
	BillingFrequency       string  `json:"billing_frequency"`
	BillingFrequencyCycles int     `json:"billing_frequency_cycles"`
	BillingStartDate       string  `json:"billing_start_date"`
}

type ProductsResponse struct {
	Success bool      `json:"success"`
	Data    []Product `json:"data"`
}

type Discount struct {
	ID          string  `json:"id"`
	Description string  `json:"description"`
	Amount      float64 `json:"amount"`
	Type        string  `json:"type"`
	DealID      int     `json:"deal_id"`
	CreatedAt   string  `json:"created_at"`
	CreatedBy   int     `json:"created_by"`
	UpdatedAt   string  `json:"updated_at"`
	UpdatedBy   int     `json:"updated_by"`
}

type DiscountsResponse struct {
	Success bool       `json:"success"`
	Data    []Discount `json:"data"`
}
