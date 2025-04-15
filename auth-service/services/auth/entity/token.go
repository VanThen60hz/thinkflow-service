package entity

type Token struct {
	Token string `json:"token"`
	// ExpiredIn in seconds
	ExpiredIn int `json:"expire_in"`
}

type TokenResponse struct {
	AccessToken Token `json:"access_token"`
	// RefreshToken will be used when access token expired
	// to issue new pair access token and refresh token.
	RefreshToken *Token `json:"refresh_token,omitempty"`
}
