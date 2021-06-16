package jwtoken

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func (t Tokens) HasAccessToken() bool {
	return t.AccessToken != ""
}

func (t Tokens) HasRefreshToken() bool {
	return t.RefreshToken != ""
}

func (t Tokens) HasBothTokens() bool {
	return t.HasAccessToken() && t.HasRefreshToken()
}
