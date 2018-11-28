package accesstoken

import (
	"testing"
	"time"

	"github.com/go-errors/errors"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2018-11-26T17:46:00Z")

	accessToken := AccessToken{
		accessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
		userID:        "userID",
		clientID:      "clientID",
		created:       now,
		expires:       now.Add(5 * time.Hour),
		scopes:        []string{},
	}

	testSubject := NewSerializer("mySuperSecret")

	serialized := testSubject.Serialize(accessToken)

	assert.Equal(t,
		"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9.0l4rKvDcFF6dfmAvtmT_LBhmxv7KAZ7FuegziUb0k7b4tmC0kWUDCtWsjUfFNAp5iSF33KvE_sNx7yNu6KLgMg",
		serialized)
}

func TestDeserialize(t *testing.T) {
	testSubject := NewSerializer("mySuperSecret")

	accessToken, err := testSubject.Deserialize("eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9.0l4rKvDcFF6dfmAvtmT_LBhmxv7KAZ7FuegziUb0k7b4tmC0kWUDCtWsjUfFNAp5iSF33KvE_sNx7yNu6KLgMg")
	assert.Nil(t, err)
	assert.NotNil(t, accessToken)

	now, _ := time.Parse(time.RFC3339, "2018-11-26T17:46:00Z")

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
	testSubject := NewSerializer("mySuperSecret")

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
