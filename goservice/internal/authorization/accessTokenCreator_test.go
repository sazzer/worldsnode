package authorization

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccessToken(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2018-11-26T17:46:00Z")
	duration, _ := time.ParseDuration("2h")

	mockClock := clock.NewMock()
	mockClock.Set(now)

	testSubject := NewAccessTokenCreator(mockClock, duration)

	accessToken := testSubject.NewAccessToken("userId", "clientId")

	expected := AccessToken{
		UserID:   "userId",
		ClientID: "clientId",
		Created:  now,
		Expires:  now.Add(duration),
		Scopes:   []string{},
	}

	assert.Equal(t, expected, accessToken)
}
