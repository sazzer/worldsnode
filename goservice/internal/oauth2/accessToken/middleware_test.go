package accesstoken

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-errors/errors"
	"github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func TestNoAuthHeader(t *testing.T) {
	serializer := testSerializer{}
	resp := performTest("", &serializer)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header().Get("Content-Type"))
	assert.Equal(t, "{\"exists\":false}", resp.Body.String())
}

func TestNoBearerAuthHeader(t *testing.T) {
	serializer := testSerializer{}
	resp := performTest("Basic abc", &serializer)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header().Get("Content-Type"))
	assert.Equal(t, "{\"exists\":false}", resp.Body.String())
}

func TestInvalidBearerAuthHeader(t *testing.T) {
	err := errors.New("Oops")
	serializer := testSerializer{"", nil, err}
	resp := performTest("Bearer abc123", &serializer)
	assert.Equal(t, http.StatusForbidden, resp.Code)
	assert.Equal(t, "application/problem+json", resp.Header().Get("Content-Type"))
	assert.Equal(t, "{\"type\":\"tag:grahamcox.co.uk,2018,worlds/problems/invalid_access_token\",\"title\":\"The provided Access Token was invalid\",\"status\":403,\"token\":\"abc123\"}", resp.Body.String())
	assert.Equal(t, "abc123", serializer.seenToken)
}

func TestVaalidBearerAuthHeader(t *testing.T) {
	accessToken := AccessToken{
		accessTokenID: "8baeefb4-dea3-4add-a4bc-b177e44b97f2",
		userID:        "userId",
		clientID:      "clientId",
		created:       now,
		expires:       now.Add(1 * time.Hour),
		scopes:        []string{},
	}

	serializer := testSerializer{"", &accessToken, nil}
	resp := performTest("Bearer abc123", &serializer)
	assert.Equal(t, http.StatusOK, resp.Code)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header().Get("Content-Type"))
	assert.Equal(t, "{\"exists\":true}", resp.Body.String())
	assert.Equal(t, "abc123", serializer.seenToken)
}

type testSerializer struct {
	seenToken   string
	accessToken *AccessToken
	error       error
}

func (t *testSerializer) Serialize(token AccessToken) string {
	return ""
}

func (t *testSerializer) Deserialize(token string) (*AccessToken, error) {
	t.seenToken = token
	return t.accessToken, t.error
}

func init() {
	gin.SetMode(gin.TestMode)
	logrus.SetLevel(logrus.DebugLevel)
}

func newRouter(serializer *testSerializer) *gin.Engine {
	router := gin.New()
	router.Use(NewMiddleware(serializer))
	router.GET("/", func(c *gin.Context) {
		_, exists := c.Get("accessToken")

		c.JSON(http.StatusOK, gin.H{
			"exists": exists,
		})
	})
	return router
}

func performTest(header string, serializer *testSerializer) *httptest.ResponseRecorder {
	router := newRouter(serializer)
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	req.Header.Set("Authorization", header)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	return w
}
