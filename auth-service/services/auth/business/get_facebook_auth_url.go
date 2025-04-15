package business

func (b *business) GetFacebookAuthURL() string {
	return b.oauthProvider.GetFacebookAuthURL()
}
