package business

func (b *business) GetGoogleAuthURL() string {
	return b.oauthProvider.GetGoogleAuthURL()
}
