package accesstoken

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-errors/errors"

	"github.com/stretchr/testify/assert"
)

type testSerializer struct {
	seenToken   string
	accessToken *AccessToken
	error       error
}

type testHandler struct {
	wasCalled bool
	req       *http.Request
}

func (t *testSerializer) Serialize(token AccessToken) string {
	return ""
}

func (t *testSerializer) Deserialize(token string) (*AccessToken, error) {
	t.seenToken = token
	return t.accessToken, t.error
}

func (t *testHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.wasCalled = true
	t.req = r
}

func TestNoAuthHeader(t *testing.T) {
	serializer := testSerializer{}
	next := testHandler{}
	req := httptest.NewRequest("GET", "/abc", nil)
	res := httptest.NewRecorder()

	accessTokenMiddleware(&serializer, &next, res, req)

	assert.True(t, next.wasCalled)
}

func TestNoBearerAuthHeader(t *testing.T) {
	serializer := testSerializer{}
	next := testHandler{}
	req := httptest.NewRequest("GET", "/abc", nil)
	req.Header.Set("Authorization", "Basic abc123")
	res := httptest.NewRecorder()

	accessTokenMiddleware(&serializer, &next, res, req)

	assert.True(t, next.wasCalled)
}

func TestInvalidBearerAuthHeader(t *testing.T) {
	err := errors.New("Oops")
	serializer := testSerializer{"", nil, err}
	next := testHandler{}
	req := httptest.NewRequest("GET", "/abc", nil)
	req.Header.Set("Authorization", "Bearer abc123")
	res := httptest.NewRecorder()

	accessTokenMiddleware(&serializer, &next, res, req)

	assert.False(t, next.wasCalled)
	assert.Equal(t, http.StatusForbidden, res.Code)
	assert.Equal(t, "abc123", serializer.seenToken)
}

func TestValidBearerAuthHeader(t *testing.T) {
	accessToken := AccessToken{
		accessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
		userID:        "userId",
		clientID:      "clientId",
		created:       now,
		expires:       now.Add(1 * time.Hour),
		scopes:        []string{},
	}

	serializer := testSerializer{"", &accessToken, nil}
	next := testHandler{}
	req := httptest.NewRequest("GET", "/abc", nil)
	req.Header.Set("Authorization", "Bearer abc123")
	res := httptest.NewRecorder()

	accessTokenMiddleware(&serializer, &next, res, req)

	assert.True(t, next.wasCalled)
	assert.Equal(t, "abc123", serializer.seenToken)
	assert.Equal(t, &accessToken, next.req.Context().Value("accessToken"))
}
