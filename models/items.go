package models

type ItemSearchResponse struct {
	Success bool `json:"success"`
	Data    struct {
		Items        []Item `json:"items"`
		RelatedItems []Item `json:"related_items"`
	} `json:"data"`
	AdditionalData `json:"additional_data"`
}

type Item struct {
	ResultScore float64 `json:"result_score"`
	Item        struct {
		ID        any    `json:"id"` // lead is string, deal, person is int
		Name      string `json:"name"`
		Type      string `json:"type"`
		Title     string `json:"title"`
		Value     int    `json:"value"`
		Currency  string `json:"currency"`
		Status    string `json:"status"`
		VisibleTo int    `json:"visible_to"`
		Owner     struct {
			ID int `json:"id"`
		} `json:"owner"`
		Stage struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"stage"`
		Person struct {
			ID   int    `json:"id"`
			Name string `json:"name"`
		} `json:"person"`
		Organization struct {
			ID      int    `json:"id"`
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"organization"`
		CustomFields []string `json:"custom_fields"`
		Notes        []string `json:"notes"`
		Phones       []string `json:"phones"`
		Emails       []any    `json:"emails"`
		PrimaryEmail any      `json:"primary_email"`
		UpdateTime   string   `json:"update_time"`
		IsArchived   bool     `json:"is_archived"`
	} `json:"item"`
}
