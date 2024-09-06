package models

type PersonFieldsResponse struct {
	Success bool          `json:"success"`
	Data    []PersonField `json:"data"`
}

type PersonFieldResponse struct {
	Success bool        `json:"success"`
	Data    PersonField `json:"data"`
}

type PersonField struct {
	ID                 int         `json:"id"`
	Key                string      `json:"key"`
	Name               string      `json:"name"`
	OrderNr            int         `json:"order_nr"`
	FieldType          string      `json:"field_type"`
	AddTime            string      `json:"add_time"`
	UpdateTime         string      `json:"update_time"`
	LastUpdatedByUser  int         `json:"last_updated_by_user_id"`
	CreatedByUser      int         `json:"created_by_user_id"`
	ActiveFlag         bool        `json:"active_flag"`
	EditFlag           bool        `json:"edit_flag"`
	IndexVisibleFlag   bool        `json:"index_visible_flag"`
	DetailsVisibleFlag bool        `json:"details_visible_flag"`
	AddVisibleFlag     bool        `json:"add_visible_flag"`
	ImportantFlag      bool        `json:"important_flag"`
	BulkEditAllowed    bool        `json:"bulk_edit_allowed"`
	SearchableFlag     bool        `json:"searchable_flag"`
	FilteringAllowed   bool        `json:"filtering_allowed"`
	SortableFlag       bool        `json:"sortable_flag"`
	Options            interface{} `json:"options,omitempty"`  // Depending on the field type, this could be a list of options
	MandatoryFlag      bool        `json:"mandatory_flag"`
}

type PersonFieldReq struct {
	Name           string      `json:"name,omitempty"`
	FieldType      string      `json:"field_type,omitempty"`
	Options        interface{} `json:"options,omitempty"` // For set/enum field types
	AddVisibleFlag bool        `json:"add_visible_flag,omitempty"`
}
