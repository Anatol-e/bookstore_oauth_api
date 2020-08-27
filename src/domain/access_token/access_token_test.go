package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.False(t, at.IsExpired(), "new access token should not be expired")
	assert.EqualValues(t, at.AccessToken, "", "new Access token should not have defined token")
	assert.True(t, at.UserId == 0, "new access token should not have an associated user")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := &AccessToken{}
	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.False(t, at.IsExpired(), "access token expired for tree hours")
}
