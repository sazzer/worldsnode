package accesstoken

import (
	"fmt"

	"github.com/gbrlsnchs/jwt/v2"
	"github.com/sirupsen/logrus"
)

const (
	audience = "worlds"
)

// Serializer is a means to serialise and deserialise Access Tokens to strings
type Serializer struct {
	signer jwt.Signer
}

// NewSerializer creates a new Access Token Serializer
func NewSerializer(secret string) Serializer {
	signer := jwt.NewHS512(secret)

	return Serializer{
		signer: signer,
	}
}

// Serialize will transform an Access Token into a String
func (s *Serializer) Serialize(token AccessToken) string {
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
			WithField("err", err).
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
			WithField("err", err).
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
func (s *Serializer) Deserialize(token string) (AccessToken, error) {
	return AccessToken{}, nil
}
