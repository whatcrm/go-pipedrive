package models

// PersonFieldsResponse is the response for all person fields.
type PersonFieldsResponse struct {
	Success bool          `json:"success"`
	Data    []PersonField `json:"data"`
}

// PersonFieldResponse is the response for a single person field.
type PersonFieldResponse struct {
	Success bool        `json:"success"`
	Data    PersonField `json:"data"`
}

// Option represents an option for set/enum type fields.
type Option struct {
	ID    int    `json:"id,omitempty"`
	Label string `json:"label"`
	Color string `json:"color,omitempty"`
}

// MonetaryField stores the numeric value and currency for monetary fields.
type MonetaryField struct {
	Value    float64 `json:"value"`
	Currency string  `json:"currency"` // ISO currency code
}

// DateRangeField represents a start date and end date for a date range field.
type DateRangeField struct {
	StartDate string `json:"start_date"` // YYYY-MM-DD
	EndDate   string `json:"end_date"`   // YYYY-MM-DD
}

// TimeRangeField represents a start time and end time for a time range field.
type TimeRangeField struct {
	StartTime string `json:"start_time"` // HH:MM:SS
	EndTime   string `json:"end_time"`   // HH:MM:SS
}

// PersonField represents the schema for a person field in the Pipedrive API.
type PersonField struct {
	ID                 int             `json:"id"`
	Key                string          `json:"key"`
	Name               string          `json:"name"`
	OrderNr            int             `json:"order_nr"`
	FieldType          string          `json:"field_type"` // Can be address, date, enum, monetary, etc.
	AddTime            string          `json:"add_time"`
	UpdateTime         string          `json:"update_time"`
	LastUpdatedByUser  int             `json:"last_updated_by_user_id"`
	CreatedByUser      int             `json:"created_by_user_id"`
	ActiveFlag         bool            `json:"active_flag"`
	EditFlag           bool            `json:"edit_flag"`
	IndexVisibleFlag   bool            `json:"index_visible_flag"`
	DetailsVisibleFlag bool            `json:"details_visible_flag"`
	AddVisibleFlag     bool            `json:"add_visible_flag"`
	ImportantFlag      bool            `json:"important_flag"`
	BulkEditAllowed    bool            `json:"bulk_edit_allowed"`
	SearchableFlag     bool            `json:"searchable_flag"`
	FilteringAllowed   bool            `json:"filtering_allowed"`
	SortableFlag       bool            `json:"sortable_flag"`
	Options            []Option        `json:"options,omitempty"`          // Options for set/enum field types
	MonetaryField      *MonetaryField  `json:"monetary_field,omitempty"`   // For monetary fields
	DateRangeField     *DateRangeField `json:"date_range_field,omitempty"` // For date-range fields
	TimeRangeField     *TimeRangeField `json:"time_range_field,omitempty"` // For time-range fields
	MandatoryFlag      bool            `json:"mandatory_flag"`
}

// PersonFieldReq represents the request body for adding/updating a person field.
type PersonFieldReq struct {
	Name           string   `json:"name,omitempty"`
	FieldType      string   `json:"field_type,omitempty"`
	Options        []Option `json:"options,omitempty"` // For set/enum field types
	AddVisibleFlag bool     `json:"add_visible_flag,omitempty"`
}
