package models

type ItemSearchResponse struct {
	Success        bool           `json:"success"`
	Data           ItemSearchData `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type ItemSearchData struct {
	Items        []Item `json:"items"`
	RelatedItems []Item `json:"related_items"`
}

type Item struct {
	ResultScore float64 `json:"result_score"`
	ItemDetails `json:"item"`
}

type ItemDetails struct {
	ID           interface{}      `json:"id"` // lead is string, others are int
	Name         string           `json:"name"`
	Type         string           `json:"type"`
	Title        string           `json:"title"`
	Value        int              `json:"value"`
	Currency     string           `json:"currency"`
	Status       string           `json:"status"`
	VisibleTo    int              `json:"visible_to"`
	Owner        ItemOwner        `json:"owner"`
	Stage        ItemStage        `json:"stage"`
	Person       ItemPerson       `json:"person"`
	Organization ItemOrganization `json:"organization"`
	CustomFields []string         `json:"custom_fields"`
	Notes        []string         `json:"notes"`
	Phones       []string         `json:"phones"`
	Emails       []interface{}    `json:"emails"`
	PrimaryEmail interface{}      `json:"primary_email"`
	UpdateTime   string           `json:"update_time"`
	IsArchived   bool             `json:"is_archived"`
}

type ItemOwner struct {
	ID int `json:"id"`
}

type ItemStage struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ItemPerson struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ItemOrganization struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
