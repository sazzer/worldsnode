package accesstoken

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccessToken(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2018-11-26T17:46:00Z")
	duration, _ := time.ParseDuration("2h")

	mockClock := clock.NewMock()
	mockClock.Set(now)

	testSubject := NewCreator(mockClock, duration)
	testSubject.idGenerator = func() uuid.UUID {
		return uuid.MustParse("8baeefb4-dea3-4add-a4bc-b177e44b97f2")
	}
	accessToken := testSubject.NewAccessToken("userId", "clientId")

	expected := AccessToken{
		AccessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
		UserID:        "userId",
		ClientID:      "clientId",
		Created:       now,
		Expires:       now.Add(duration),
		Scopes:        []string{},
	}

	assert.Equal(t, expected, accessToken)
}
