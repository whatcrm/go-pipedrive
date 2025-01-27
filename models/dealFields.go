package models

// DealFieldsResponse represents the response for all deal fields.
type DealFieldsResponse struct {
	Success        bool        `json:"success"`
	Data           []DealField `json:"data"`
	AdditionalData Pagination  `json:"additional_data"`
}

// DealFieldResponse represents the response for a single deal field.
type DealFieldResponse struct {
	Success bool      `json:"success"`
	Data    DealField `json:"data"`
}

// ShowInPipelines represents the pipelines where the deal field is visible.
type ShowInPipelines struct {
	ShowInAll   bool  `json:"show_in_all"`
	PipelineIDs []int `json:"pipeline_ids"`
}

// Option represents an option for set/enum type fields.
type DealOption struct {
	ID    any    `json:"id,omitempty"`
	Label string `json:"label"`
	AltID string `json:"alt_id,omitempty"` // Alternate ID for options (e.g., Prospector)
}

// DealField represents a single deal field in the Pipedrive API.
type DealField struct {
	ID                        int             `json:"id,omitempty"`
	Key                       string          `json:"key"`
	Name                      string          `json:"name"`
	GroupID                   *int            `json:"group_id,omitempty"`
	OrderNr                   int             `json:"order_nr"`
	FieldType                 string          `json:"field_type"`
	JSONColumnFlag            bool            `json:"json_column_flag"`
	AddTime                   string          `json:"add_time"`
	UpdateTime                *string         `json:"update_time,omitempty"`
	LastUpdatedByUserID       *int            `json:"last_updated_by_user_id,omitempty"`
	EditFlag                  bool            `json:"edit_flag"`
	DetailsVisibleFlag        bool            `json:"details_visible_flag"`
	AddVisibleFlag            bool            `json:"add_visible_flag"`
	ImportantFlag             bool            `json:"important_flag"`
	BulkEditAllowed           bool            `json:"bulk_edit_allowed"`
	FilteringAllowed          bool            `json:"filtering_allowed"`
	SortableFlag              bool            `json:"sortable_flag"`
	MandatoryFlag             interface{}     `json:"mandatory_flag"` // This can be bool or map for conditional flags
	SearchableFlag            bool            `json:"searchable_flag"`
	Description               *string         `json:"description,omitempty"`
	CreatedByUserID           *int            `json:"created_by_user_id,omitempty"`
	ActiveFlag                bool            `json:"active_flag"`
	ProjectsDetailVisibleFlag bool            `json:"projects_detail_visible_flag"`
	ShowInPipelines           ShowInPipelines `json:"show_in_pipelines"`
	Options                   []DealOption    `json:"options,omitempty"`
	ParentID                  *int            `json:"parent_id,omitempty"`
	IDSuffix                  *string         `json:"id_suffix,omitempty"`
	IsSubfield                bool            `json:"is_subfield,omitempty"`
	UseField                  *string         `json:"use_field,omitempty"` // Only for specific field types like `stage`
	Link                      *string         `json:"link,omitempty"`      // Used for certain fields like `deal` links
}

// DealFieldReq represents the request body for adding/updating a deal field.
type DealFieldReq struct {
	Name           string       `json:"name,omitempty"`
	FieldType      string       `json:"field_type,omitempty"`
	Options        []DealOption `json:"options,omitempty"` // For set/enum field types
	AddVisibleFlag bool         `json:"add_visible_flag,omitempty"`
}
