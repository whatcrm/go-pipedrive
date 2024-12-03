package models

type Product struct {
	ID               int     `json:"id"`
	Name             string  `json:"name"`
	Code             string  `json:"code,omitempty"`
	Description      string  `json:"description,omitempty"`
	Unit             string  `json:"unit,omitempty"`
	Tax              float64 `json:"tax,omitempty"`
	ActiveFlag       bool    `json:"active_flag"`
	Selectable       bool    `json:"selectable"`
	VisibleTo        string  `json:"visible_to"`
	BillingFrequency string  `json:"billing_frequency,omitempty"`
}

type ProductRequest struct {
	Name             string  `json:"name"`
	Code             string  `json:"code,omitempty"`
	Description      string  `json:"description,omitempty"`
	Unit             string  `json:"unit,omitempty"`
	Tax              float64 `json:"tax,omitempty"`
	ActiveFlag       bool    `json:"active_flag"`
	Selectable       bool    `json:"selectable"`
	VisibleTo        string  `json:"visible_to,omitempty"`
	OwnerID          int     `json:"owner_id,omitempty"`
	Prices           []Price `json:"prices,omitempty"`
	BillingFrequency string  `json:"billing_frequency,omitempty"`
}

type ProductUpdateRequest struct {
	Name             string  `json:"name,omitempty"`
	Code             string  `json:"code,omitempty"`
	Description      string  `json:"description,omitempty"`
	Unit             string  `json:"unit,omitempty"`
	Tax              float64 `json:"tax,omitempty"`
	ActiveFlag       bool    `json:"active_flag"`
	Selectable       bool    `json:"selectable"`
	VisibleTo        string  `json:"visible_to,omitempty"`
	Prices           []Price `json:"prices,omitempty"`
	BillingFrequency string  `json:"billing_frequency,omitempty"`
}

type ProductResponse struct {
	Success bool    `json:"success"`
	Data    Product `json:"data"`
}

type ProductsResponse struct {
	Success bool      `json:"success"`
	Data    []Product `json:"data"`
}

type Price struct {
	Currency      string  `json:"currency"`
	Price         float64 `json:"price"`
	Cost          float64 `json:"cost,omitempty"`
	OverheadCost  float64 `json:"overhead_cost,omitempty"`
	Notes         string  `json:"notes,omitempty"`
}

type ProductVariation struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Currency string  `json:"currency"`
	Cost     float64 `json:"cost,omitempty"`
}

type ProductVariationRequest struct {
	Name   string  `json:"name"`
	Prices []Price `json:"prices"`
}

type ProductVariationResponse struct {
	Success bool             `json:"success"`
	Data    ProductVariation `json:"data"`
}
