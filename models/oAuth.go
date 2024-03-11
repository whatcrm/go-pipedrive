package models

type AccessTokenRequest struct {
	GrantType   string `json:"grant_type"`
	Code        string `json:"code"`
	RedirectURI string `json:"redirect_uri"`
}

type RefreshTokenRequest struct {
	GrantType   string `json:"grant_type"`
	RefreshToken string `json:"refresh_token"`
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	APIDomain    string `json:"api_domain"`
}