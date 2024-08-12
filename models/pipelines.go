package models

type Pipeline struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	URLTitle        string `json:"url_title"`
	OrderNr         int    `json:"order_nr"`
	Active          bool   `json:"active"`
	DealProbability bool   `json:"deal_probability"`
	AddTime         string `json:"add_time"`
	UpdateTime      string `json:"update_time"`
	Selected        bool   `json:"selected"`
}

type PipelinesResponse struct {
	Success        bool       `json:"success"`
	Data           []Pipeline `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}

type PipelineResponse struct {
	Success        bool     `json:"success"`
	Data           Pipeline `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}

type Stage struct {
	ID                      int    `json:"id"`
	OrderNr                 int    `json:"order_nr"`
	Name                    string `json:"name"`
	ActiveFlag              bool   `json:"active_flag"`
	DealProbability         int    `json:"deal_probability"`
	PipelineID              int    `json:"pipeline_id"`
	RottenFlag              bool   `json:"rotten_flag"`
	RottenDays              int    `json:"rotten_days"`
	AddTime                 string `json:"add_time"`
	UpdateTime              string `json:"update_time"`
	PipelineName            string `json:"pipeline_name"`
	PipelineDealProbability bool   `json:"pipeline_deal_probability"`
}

type StagesResponse struct {
	Success        bool    `json:"success"`
	Data           []Stage `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}

type StageResponse struct {
	Success        bool  `json:"success"`
	Data           Stage `json:"data"`
	AdditionalData struct {
		Pagination struct {
			Start                 int  `json:"start"`
			Limit                 int  `json:"limit"`
			MoreItemsInCollection bool `json:"more_items_in_collection"`
		} `json:"pagination"`
	} `json:"additional_data"`
}
