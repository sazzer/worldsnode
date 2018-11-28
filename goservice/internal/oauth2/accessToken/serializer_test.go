package accesstoken

import (
	"testing"
	"time"

	"github.com/benbjohnson/clock"
	"github.com/go-errors/errors"

	"github.com/stretchr/testify/assert"
)

// currentTime is the time to use as the current system time in the tests
var currentTime = "2018-11-26T17:46:00Z"

// now is a parsed version of currentTime
var now, _ = time.Parse(time.RFC3339, currentTime)

// accessToken is the Access Token to work with
var accessToken = AccessToken{
	accessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
	userID:        "userID",
	clientID:      "clientID",
	created:       now,
	expires:       now.Add(5 * time.Hour),
	scopes:        []string{},
}

// secret is the secret used to sign serializedAccessToken
var secret = "mySuperSecret"

// serialziedAccessToken is a previously computed JWT that matches accessToken
var serializedAccessToken = "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9.0l4rKvDcFF6dfmAvtmT_LBhmxv7KAZ7FuegziUb0k7b4tmC0kWUDCtWsjUfFNAp5iSF33KvE_sNx7yNu6KLgMg"

func TestSerialize(t *testing.T) {

	mockClock := clock.NewMock()
	testSubject := NewSerializer(secret, mockClock)

	serialized := testSubject.Serialize(accessToken)

	assert.Equal(t, serializedAccessToken, serialized)
}

func TestDeserialize(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, currentTime)
	mockClock := clock.NewMock()
	mockClock.Set(now)

	testSubject := NewSerializer(secret, mockClock)

	accessToken, err := testSubject.Deserialize("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9.0l4rKvDcFF6dfmAvtmT_LBhmxv7KAZ7FuegziUb0k7b4tmC0kWUDCtWsjUfFNAp5iSF33KvE_sNx7yNu6KLgMg")
	assert.Nil(t, err)
	assert.NotNil(t, accessToken)

	expected := AccessToken{
		accessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
		userID:        "userID",
		clientID:      "clientID",
		created:       now,
		expires:       now.Add(5 * time.Hour),
		scopes:        []string{},
	}

	assert.Equal(t, &expected, accessToken)
}

func TestDeserializeErrors(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, currentTime)
	mockClock := clock.NewMock()
	mockClock.Set(now)

	testSubject := NewSerializer(secret, mockClock)

	tests := []string{
		"",
		"abc",
		"a.b.c",
		"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9.0l4rKvDcFF6dfmAvtmT_LBhmxv7KAZ7FuegziUb0k7b4tmC0kWUDCtWsjUfFNAp5iSF33KvE_sNx7yNu6KLgMf",
		"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9",
	}

	for _, tt := range tests {
		accessToken, err := testSubject.Deserialize(tt)

		assert.NotNil(t, err)
		assert.Nil(t, accessToken)

		assert.True(t, errors.Is(err, InvalidAccessTokenError))
	}
}

func TestDeserializeTimeErrors(t *testing.T) {
	mockClock := clock.NewMock()

	testSubject := NewSerializer(secret, mockClock)

	tests := []string{
		"2018-11-26T17:45:00Z", // A minute before the start
		"2018-11-26T22:47:00Z", // A minute after the end
	}

	for _, tt := range tests {
		now, _ := time.Parse(time.RFC3339, tt)
		mockClock.Set(now)

		accessToken, err := testSubject.Deserialize(serializedAccessToken)

		assert.NotNil(t, err)
		assert.Nil(t, accessToken)

		assert.True(t, errors.Is(err, InvalidAccessTokenError))
	}
}
