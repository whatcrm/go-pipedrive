package models

type UsersResponse struct {
	Success        bool           `json:"success"`
	Data           []User         `json:"data"`
	AdditionalData AdditionalData `json:"additional_data"`
}

type User struct {
	ID                int      `json:"id"`
	Name              string   `json:"name"`
	Email             string   `json:"email"`
	Lang              int      `json:"lang"`
	Locale            string   `json:"locale,omitempty"`
	TimezoneName      string   `json:"timezone_name,omitempty"`
	TimezoneOffset    string   `json:"timezone_offset,omitempty"`
	DefaultCurrency   string   `json:"default_currency,omitempty"`
	IconURL           string   `json:"icon_url,omitempty"`
	ActiveFlag        bool     `json:"active_flag"`
	IsAdmin           int      `json:"is_admin"`
	RoleID            int      `json:"role_id"`
	Created           string   `json:"created"`
	HasCreatedCompany bool     `json:"has_created_company"`
	IsYou             bool     `json:"is_you"`
	Access            []Access `json:"access"`
	LastLogin         *string   `json:"last_login,omitempty"`
	Modified          string   `json:"modified"`
	Phone             *string   `json:"phone,omitempty"`
	IsDeleted         bool     `json:"is_deleted"`
}

type Access struct {
	App             string `json:"app"`
	Admin           bool   `json:"admin"`
	PermissionSetID string `json:"permission_set_id"`
}

type UserResponse struct {
	Success bool     `json:"success"`
	Data    UserData `json:"data"`
}

type UserData struct {
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
}

type Language struct {
	LanguageCode string `json:"language_code"`
	CountryCode  string `json:"country_code"`
}
