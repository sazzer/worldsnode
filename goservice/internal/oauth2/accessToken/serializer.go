package accesstoken

import (
	"errors"
	"fmt"
	"time"

	"github.com/benbjohnson/clock"

	"grahamcox.co.uk/worlds/service/internal/oauth2/clients"
	"grahamcox.co.uk/worlds/service/internal/users"

	"github.com/gbrlsnchs/jwt/v2"
	"github.com/sirupsen/logrus"
)

const (
	audience = "worlds"
)

// ErrInvalidAccessToken is used to indicate that an access token we are deserializing was invalid
var ErrInvalidAccessToken = errors.New("Invalid Access Token")

// Serializer is a means to serialise and deserialise Access Tokens to strings
type Serializer interface {
	Serialize(token AccessToken) string
	Deserialize(token string) (*AccessToken, error)
}

// serializerImpl is a means to serialise and deserialise Access Tokens to strings
type serializerImpl struct {
	signer jwt.Signer
	clock  clock.Clock
}

// NewSerializer creates a new Access Token Serializer
func NewSerializer(secret string, clock clock.Clock) Serializer {
	signer := jwt.NewHS512(secret)

	return &serializerImpl{
		signer: signer,
		clock:  clock,
	}
}

// Serialize will transform an Access Token into a String
func (s *serializerImpl) Serialize(token AccessToken) string {
	logrus.WithField("accessToken", token).Debug("Serializing access token")

	jot := &jwt.JWT{
		Issuer:         string(token.ClientID()),
		Subject:        string(token.UserID()),
		Audience:       audience,
		ExpirationTime: token.Expires().Unix(),
		NotBefore:      token.Created().Unix(),
		IssuedAt:       token.Created().Unix(),
		ID:             string(token.ID()),
	}
	jot.SetAlgorithm(s.signer)

	payload, err := jwt.Marshal(jot)
	if err != nil {
		logrus.
			WithField("accessToken", token).
			WithField("jot", jot).
			WithError(err).
			Warn("Error creating JWT Payload")
	}

	logrus.
		WithField("accessToken", token).
		WithField("jot", jot).
		Debug("Created JWT Payload")

	signed, err := s.signer.Sign(payload)
	if err != nil {
		logrus.
			WithField("accessToken", token).
			WithField("jot", jot).
			WithError(err).
			Warn("Error signing JWT Payload")
	}

	encoded := fmt.Sprintf("%s", signed)

	logrus.
		WithField("accessToken", token).
		WithField("jot", jot).
		WithField("signed", encoded).
		Debug("Signed JWT Payload")

	return encoded
}

// Deserialize will transform a String into an Access Token, or else return an error if the string
// wasn't a valid serialized Access Token created by the Serialize() function
func (s *serializerImpl) Deserialize(token string) (*AccessToken, error) {
	payload, sig, err := jwt.Parse(token)
	if err != nil {
		logrus.
			WithField("accessToken", token).
			WithError(err).
			Warn("Error parsing access token")
		return nil, ErrInvalidAccessToken
	}

	if err = s.signer.Verify(payload, sig); err != nil {
		logrus.
			WithField("accessToken", token).
			WithError(err).
			Warn("Error verifying access token")
		return nil, ErrInvalidAccessToken
	}

	var jot jwt.JWT
	if err = jwt.Unmarshal(payload, &jot); err != nil {
		logrus.
			WithField("accessToken", token).
			WithError(err).
			Warn("Error unmarshalling access token")
		return nil, ErrInvalidAccessToken
	}

	now := s.clock.Now()
	audienceValidator := jwt.AudienceValidator(audience)
	issuedAtValidator := jwt.IssuedAtValidator(now)
	expiresValidator := jwt.ExpirationTimeValidator(now)

	if err = jot.Validate(audienceValidator, issuedAtValidator, expiresValidator); err != nil {
		logrus.
			WithField("accessToken", token).
			WithField("jot", jot).
			WithError(err).
			Warn("Error validating access token")
		return nil, ErrInvalidAccessToken
	}

	result := AccessToken{
		accessTokenID: ID(jot.ID),
		userID:        users.ID(jot.Subject),
		clientID:      clients.ID(jot.Issuer),
		created:       time.Unix(jot.IssuedAt, 0).UTC(),
		expires:       time.Unix(jot.ExpirationTime, 0).UTC(),
		scopes:        []string{},
	}

	logrus.
		WithField("accessToken", token).
		WithField("result", result).
		Info("Parsed access token")

	return &result, nil
}
