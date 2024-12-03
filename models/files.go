package models

import "fmt"

type File struct {
	ID             int    `json:"id"`
	UserID         int    `json:"user_id"`
	DealID         int    `json:"deal_id"`
	PersonID       int    `json:"person_id"`
	OrgID          int    `json:"org_id"`
	ProductID      int    `json:"product_id"`
	ActivityID     int    `json:"activity_id"`
	LeadID         string `json:"lead_id"`
	LogID          int    `json:"log_id"`
	AddTime        string `json:"add_time"`
	UpdateTime     string `json:"update_time"`
	FileName       string `json:"file_name"`
	FileType       string `json:"file_type"`
	FileSize       int    `json:"file_size"`
	ActiveFlag     bool   `json:"active_flag"`
	InlineFlag     bool   `json:"inline_flag"`
	RemoteLocation string `json:"remote_location"`
	RemoteID       string `json:"remote_id"`
	CID            string `json:"cid"`
	S3Bucket       string `json:"s3_bucket"`
	MailMessageID  string `json:"mail_message_id"`
	MailTemplateID string `json:"mail_template_id"`
	DealName       string `json:"deal_name"`
	PersonName     string `json:"person_name"`
	OrgName        string `json:"org_name"`
	ProductName    string `json:"product_name"`
	LeadName       string `json:"lead_name"`
	URL            string `json:"url"`
	Name           string `json:"name"`
	Description    string `json:"description"`
}

type FileRequest struct {
	DealID      int    `json:"deal_id,omitempty"`
	PersonID    int    `json:"person_id,omitempty"`
	OrgID       int    `json:"org_id,omitempty"`
	ProductID   int    `json:"product_id,omitempty"`
	ActivityID  int    `json:"activity_id,omitempty"`
	LeadID      string `json:"lead_id,omitempty"`
	Description string `json:"description,omitempty"`
}

type RemoteFileRequest struct {
	FileType       string `json:"file_type"`
	Title          string `json:"title"`
	ItemType       string `json:"item_type"`
	ItemID         int    `json:"item_id"`
	RemoteLocation string `json:"remote_location"`
}

type RemoteLinkRequest struct {
	ItemType       string `json:"item_type"`
	ItemID         int    `json:"item_id"`
	RemoteID       string `json:"remote_id"`
	RemoteLocation string `json:"remote_location"`
}

type FileUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
}

type FileResponse struct {
	Success bool `json:"success"`
	Data    File `json:"data"`
}

type FilesResponse struct {
	Success               bool   `json:"success"`
	Data                  []File `json:"data"`
	MoreItemsInCollection bool   `json:"more_items_in_collection"`
	NextStart             int    `json:"next_start"`
}

type FileDownloadResponse struct {
	Success bool `json:"success"`
	Data    struct {
		URL string `json:"url"`
	} `json:"data"`
}

func (f FileRequest) ToMap() map[string]string {
	result := make(map[string]string)
	if f.DealID != 0 {
		result["deal_id"] = fmt.Sprintf("%d", f.DealID)
	}
	if f.PersonID != 0 {
		result["person_id"] = fmt.Sprintf("%d", f.PersonID)
	}
	if f.OrgID != 0 {
		result["org_id"] = fmt.Sprintf("%d", f.OrgID)
	}
	if f.ProductID != 0 {
		result["product_id"] = fmt.Sprintf("%d", f.ProductID)
	}
	if f.ActivityID != 0 {
		result["activity_id"] = fmt.Sprintf("%d", f.ActivityID)
	}
	if f.LeadID != "" {
		result["lead_id"] = f.LeadID
	}
	if f.Description != "" {
		result["description"] = f.Description
	}
	return result
}
