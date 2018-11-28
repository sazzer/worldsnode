package accesstoken

import (
	"context"
	"net/http"
	"strings"

	"github.com/sirupsen/logrus"
)

// accessTokenMiddleware is the actual middleware that will do the processing.
// If there is no Authorization header then the request is passed on
// If there is an Authorization header but it is not a "Bearer " prefix then the request is passed on
// The bearer token is then parsed using the serializer, and if parsing fails then an HTTP 403 error is returned
// If parsing succeeds then the access token is added to the request context for later use
func accessTokenMiddleware(serializer Serializer, next http.Handler, w http.ResponseWriter, r *http.Request) {
	authorizationHeader := r.Header.Get("Authorization")

	logrus.
		WithField("authorization", authorizationHeader).
		Debug("Processing authorization header")

	if authorizationHeader == "" {
		logrus.
			WithField("authorization", authorizationHeader).
			Debug("No authorization header present")
		next.ServeHTTP(w, r)
		return
	}

	if !strings.HasPrefix(authorizationHeader, "Bearer ") {
		logrus.
			WithField("authorization", authorizationHeader).
			Debug("Authorization header present is not a Bearer Token")
		next.ServeHTTP(w, r)
		return
	}

	token := authorizationHeader[7:]
	parsed, err := serializer.Deserialize(token)
	if err != nil {
		logrus.
			WithError(err).
			WithField("authorization", authorizationHeader).
			Warn("Failed to parse access token")
		w.WriteHeader(http.StatusForbidden)
	} else {
		logrus.
			WithField("authorization", authorizationHeader).
			WithField("parsed", parsed).
			Warn("Parsed access token")

		ctx := context.WithValue(r.Context(), "accessToken", parsed)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// NewMiddleware will create some HTTP Middleware that will parse the Access Token from the incoming request, if there is one
// and either return an error or allow processing to continue as appropriate
func NewMiddleware(serializer Serializer) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			accessTokenMiddleware(serializer, next, w, r)
		})
	}
}
