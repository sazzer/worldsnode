package accesstoken

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// accessTokenMiddleware is the actual middleware that will do the processing.
// If there is no Authorization header then the request is passed on
// If there is an Authorization header but it is not a "Bearer " prefix then the request is passed on
// The bearer token is then parsed using the serializer, and if parsing fails then an HTTP 403 error is returned
// If parsing succeeds then the access token is added to the request context for later use
func accessTokenMiddleware(serializer Serializer, c *gin.Context) {
	authorizationHeader := c.Request.Header.Get("Authorization")

	logrus.
		WithField("authorization", authorizationHeader).
		Debug("Processing authorization header")

	if authorizationHeader == "" {
		logrus.
			WithField("authorization", authorizationHeader).
			Debug("No authorization header present")
		c.Next()
		return
	}

	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		logrus.
			WithField("authorization", authorizationHeader).
			Debug("Authorization header present is not a Bearer Token")
		c.Next()
		return
	}

	token := authorizationHeader[7:]
	parsed, err := serializer.Deserialize(token)
	if err != nil {
		logrus.
			WithError(err).
			WithField("authorization", authorizationHeader).
			Warn("Failed to parse access token")
		c.AbortWithStatus(http.StatusForbidden)
	} else {
		logrus.
			WithField("authorization", authorizationHeader).
			WithField("parsed", parsed).
			Debug("Parsed access token")

		c.Set("accessToken", parsed)
		c.Next()
	}
}

// NewMiddleware will create some HTTP Middleware that will parse the Access Token from the incoming request, if there is one
// and either return an error or allow processing to continue as appropriate
func NewMiddleware(serializer Serializer) func(*gin.Context) {
	return func(c *gin.Context) {
		accessTokenMiddleware(serializer, c)
	}
}
