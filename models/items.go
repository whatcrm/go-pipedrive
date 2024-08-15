package models

type ItemSearchResponse struct {
	Success        bool   `json:"success"`
	Data           []Item `json:"data"`
	RelatedItems   []Item `json:"related_items"`
	AdditionalData `json:"additional_data"`
}
type Item struct {
	ResultScore float64 `json:"result_score"`
	Item        struct {
		ID        int    `json:"id"`
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
	} `json:"item"`
}
