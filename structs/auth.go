package structs

type AuthResponse struct {
	AccessToken          string `json:"access_token"`
	AccessTokenExpiresIn uint   `json:"access_token_expires_in"`
	OrganizationName     string `json:"organization_name"`
	DeveloperEmail       string `json:"developer_email"`
	TokenType            string `json:"token_type"`
}
