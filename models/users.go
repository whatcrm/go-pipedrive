package models

type UsersResponse struct {
	Success        bool           `json:"success"`
	Data           []User         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type Access struct {
	App             string `json:"app"`
	Admin           bool   `json:"admin"`
	PermissionSetID string `json:"permission_set_id"`
}

type UserResponse struct {
	Success bool     `json:"success"`
	Data    User `json:"data"`
}

type User struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Lang              int      `json:"lang"`
	Locale            string   `json:"locale"`
	TimezoneName      string   `json:"timezone_name"`
	TimezoneOffset    string   `json:"timezone_offset"`
	DefaultCurrency   string   `json:"default_currency"`
	IconURL           *string  `json:"icon_url"` 
	ActiveFlag        bool     `json:"active_flag"`
	IsAdmin           int      `json:"is_admin"`
	RoleID            int      `json:"role_id"`
	Created           string   `json:"created"`
	HasCreatedCompany bool     `json:"has_created_company"`
	IsYou             bool     `json:"is_you"`
	Access            []Access `json:"access"`
	Phone             *string  `json:"phone"` 
	LastLogin         *string  `json:"last_login"`
	Modified          string   `json:"modified"`
	IsDeleted         bool     `json:"is_deleted"`
	CompanyID         int      `json:"company_id"`
	CompanyName       string   `json:"company_name"`
	CompanyDomain     string   `json:"company_domain"`
	CompanyCountry    string   `json:"company_country"`
	Language          Language `json:"language"`
	Activated         bool     `json:"activated"`
	CompanyIndustry   *string  `json:"company_industry,omitempty"`
}

type Language struct {
	LanguageCode string `json:"language_code"`
	CountryCode  string `json:"country_code"`
}
