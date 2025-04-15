package business

import (
	"thinkflow-service/common"

	"github.com/VanThen60hz/service-context/core"
)

func (b *business) ValidateOAuthState(receivedState string) error {
	if receivedState != common.OAuthStateString {
		return core.ErrBadRequest.WithError("invalid oauth state")
	}
	return nil
}
