package models

// Person represents a person in the Pipedrive system.
type Person struct {
	ID         int      `json:"id"`
	ActiveFlag bool     `json:"active_flag"`
	Name       string   `json:"name"`
	Emails     []Emails `json:"emails"`
	Phones     []Phones `json:"phones"`
	OwnerID    int      `json:"owner_id"`
	OrgID      int      `json:"org_id"`
}

// Person represents a person in the Pipedrive system.
type PersonsRequest struct {
	Name         string                 `json:"name"`
	OwnerID      int                    `json:"owner_id,omitempty"`
	OrgID        int                    `json:"org_id,omitempty"`
	Emails       []string               `json:"emails,omitempty"`
	Phones       []string               `json:"phones,omitempty"`
	VisibleTo    int                    `json:"visible_to,omitempty"`
	AddTime      string                 `json:"add_time,omitempty"`
	CustomFields map[string]interface{} `json:"custom_fields,omitempty"`
	Label        int                    `json:"label,omitempty"`
	LabelIDs     []int                  `json:"label_ids,omitempty"`
}

type PersonsResponse struct {
	Success        bool     `json:"success"`
	Data           []Person `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}

type PersonResponse struct {
	Success        bool   `json:"success"`
	Data           Person `json:"data"`
	AdditionalData struct {
		NextCursor string `json:"next_cursor"`
	} `json:"additional_data"`
}

// Email represents an email in the Pipedrive system.
type Emails struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// Phone represents a phone number in the Pipedrive system.
type Phones struct {
	Value   string `json:"value"`
	Primary bool   `json:"primary"`
}

// PersonPictureRequest holds the data for adding a picture to a person
type PersonPictureRequest struct {
	FileName   string //required
	FilePath   string //required The URL or file path of the image
	CropX      int    // X coordinate for cropping
	CropY      int    // Y coordinate for cropping
	CropWidth  int    // Width of the cropping area
	CropHeight int    // Height of the cropping area
}

// PersonPictureResponse represents the response when uploading a picture to a person.
type PersonPictureResponse struct {
	Success       bool       `json:"success"`
	ID            int        `json:"id"`
	ItemType      string     `json:"item_type"`
	ItemID        int        `json:"item_id"`
	ActiveFlag    bool       `json:"active_flag"`
	AddTime       string     `json:"add_time"`
	UpdateTime    string     `json:"update_time"`
	AddedByUserID int        `json:"added_by_user_id"`
	Images        ImageLinks `json:"images"`
}

// ImageLinks holds the links for the different sizes of the uploaded picture.
type ImageLinks struct {
	Image128 string `json:"128"`
	Image512 string `json:"512"`
}
