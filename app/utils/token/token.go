package token

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
	"time"
)

type Claims struct {
	jwt.Claims
}

func (c *Claims) SignHS256(hmacKey []byte) (string, error) {
	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.HS256, Key: hmacKey},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		return "", errors.Wrap(err, "NewSigner")
	}
	return jwt.Signed(signer).Claims(c).CompactSerialize()
}

func Parse(tokenStr string, key []byte) (*Claims, error) {
	token, err := jwt.ParseSigned(tokenStr)
	if err != nil {
		return nil, errors.Wrap(err, "ParseSigned")
	}

	claims := Claims{}
	err = token.Claims(key, &claims)
	if err != nil {
		return nil, errors.Wrap(err, "Claims")
	}

	err = claims.Claims.Validate(jwt.Expected{
		Audience: jwt.Audience{"github.com/dkeohane/yagsy"}, // jwt.Audience{cfg.AuthNURL.String()},
		Issuer:   "yagsy",                                   // cfg.AuthNURL.String(),
	})
	if err != nil {
		return nil, errors.Wrap(err, "Validate")
	}
	/*
		if claims.Scope != scope {
			return nil, fmt.Errorf("token scope not valid")
		}*/

	return &claims, nil
}

func New(userID uuid.UUID, audience string) *Claims {
	return &Claims{
		Claims: jwt.Claims{
			Issuer:   "yagsy",
			Subject:  userID.String(),
			Audience: jwt.Audience{audience},
			Expiry:   jwt.NewNumericDate(time.Now().Add(time.Hour)), // Expire in 1 hour.
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}
}
