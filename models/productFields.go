package models

type ProductField struct {
	ID            int           `json:"id"`
	Key           string        `json:"key"`
	Name          string        `json:"name"`
	FieldType     string        `json:"field_type"`
	AddTime       string        `json:"add_time"`
	UpdateTime    string        `json:"update_time"`
	IsDeleted     bool          `json:"is_deleted"`
	EditFlag      bool          `json:"edit_flag"`
	MandatoryFlag bool          `json:"mandatory_flag"`
	Options       []FieldOption `json:"options,omitempty"`
}

type ProductFieldRequest struct {
	Name      string        `json:"name"`
	FieldType string        `json:"field_type"`
	Options   []FieldOption `json:"options,omitempty"`
}

type ProductFieldUpdateRequest struct {
	Name    string        `json:"name,omitempty"`
	Options []FieldOption `json:"options,omitempty"`
}

type ProductFieldResponse struct {
	Success bool         `json:"success"`
	Data    ProductField `json:"data"`
}

type ProductFieldsResponse struct {
	Success bool           `json:"success"`
	Data    []ProductField `json:"data"`
}
