package models

// NoteField represents a single note field.
type NoteField struct {
	ID               int    `json:"id"`
	Key              string `json:"key"`
	Name             string `json:"name"`
	FieldType        string `json:"field_type"`
	ActiveFlag       bool   `json:"active_flag"`
	EditFlag         bool   `json:"edit_flag"`
	BulkEditAllowed  bool   `json:"bulk_edit_allowed"`
	MandatoryFlag    bool   `json:"mandatory_flag"`
}

// NoteFieldsResponse represents the response for getting all note fields.
type NoteFieldsResponse struct {
	Success bool        `json:"success"`
	Data    []NoteField `json:"data"`
}
