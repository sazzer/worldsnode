package accesstoken

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSerialize(t *testing.T) {
	now, _ := time.Parse(time.RFC3339, "2018-11-26T17:46:00Z")

	accessToken := AccessToken{
		AccessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
		UserID:        "userID",
		ClientID:      "clientID",
		Created:       now,
		Expires:       now.Add(5 * time.Hour),
		Scopes:        []string{},
	}

	testSubject := NewSerializer("mySuperSecret")

	serialized := testSubject.Serialize(accessToken)

	assert.Equal(t,
		"eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjbGllbnRJRCIsInN1YiI6InVzZXJJRCIsImF1ZCI6IndvcmxkcyIsImV4cCI6MTU0MzI3MjM2MCwibmJmIjoxNTQzMjU0MzYwLCJpYXQiOjE1NDMyNTQzNjAsImp0aSI6IjhiYWVlZmI0LWRlYTMtNGFkZC1hNGJjLWIxNzdlNDRiOTdmMiJ9.0l4rKvDcFF6dfmAvtmT_LBhmxv7KAZ7FuegziUb0k7b4tmC0kWUDCtWsjUfFNAp5iSF33KvE_sNx7yNu6KLgMg",
		serialized)
}
