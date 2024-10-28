package models

type ActivityField struct {
	ID                    int      `json:"id"`
	Key                   string   `json:"key"`
	Name                  string   `json:"name"`
	OrderNr               int      `json:"order_nr"`
	FieldType             string   `json:"field_type"`
	AddTime               string   `json:"add_time"`
	UpdateTime            string   `json:"update_time"`
	LastUpdatedByUserID   int      `json:"last_updated_by_user_id"`
	CreatedByUserID       int      `json:"created_by_user_id"`
	ActiveFlag            bool     `json:"active_flag"`
	EditFlag              bool     `json:"edit_flag"`
	IndexVisibleFlag      bool     `json:"index_visible_flag"`
	DetailsVisibleFlag    bool     `json:"details_visible_flag"`
	AddVisibleFlag        bool     `json:"add_visible_flag"`
	ImportantFlag         bool     `json:"important_flag"`
	BulkEditAllowed       bool     `json:"bulk_edit_allowed"`
	SearchableFlag        bool     `json:"searchable_flag"`
	FilteringAllowed      bool     `json:"filtering_allowed"`
	SortableFlag          bool     `json:"sortable_flag"`
	Options               []Option `json:"options,omitempty"`
	MandatoryFlag         bool     `json:"mandatory_flag"`
}

type ActivityFieldsResponse struct {
	Success bool           `json:"success"`
	Data    []ActivityField `json:"data"`
}
