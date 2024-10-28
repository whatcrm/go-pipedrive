package models

type Currency struct {
	ID            int    `json:"id"`
	Code          string `json:"code"`
	Name          string `json:"name"`
	DecimalPoints int    `json:"decimal_points"`
	Symbol        string `json:"symbol"`
	ActiveFlag    bool   `json:"active_flag"`
	IsCustomFlag  bool   `json:"is_custom_flag"`
}

type CurrenciesResponse struct {
	Success bool       `json:"success"`
	Data    []Currency `json:"data"`
}
