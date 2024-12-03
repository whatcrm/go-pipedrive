package models

type ActivityType struct {
	ID           int    `json:"id"`
	OrderNr      int    `json:"order_nr"`
	Name         string `json:"name"`
	KeyString    string `json:"key_string"`
	IconKey      string `json:"icon_key"`
	ActiveFlag   bool   `json:"active_flag"`
	Color        string `json:"color"`
	IsCustomFlag bool   `json:"is_custom_flag"`
	AddTime      string `json:"add_time"`
	UpdateTime   string `json:"update_time"`
}

type ActivityTypesResponse struct {
	Success bool           `json:"success"`
	Data    []ActivityType `json:"data"`
}

type ActivityTypeResponse struct {
	Success bool         `json:"success"`
	Data    ActivityType `json:"data"`
}

type ActivityTypeRequest struct {
	Name     string `json:"name"`
	IconKey  string `json:"icon_key"`
	Color    string `json:"color,omitempty"`
	OrderNr  int    `json:"order_nr,omitempty"`
}
