package auth

import "github.com/pkg/errors"

var (
	ErrUnexpectedSigningMethod = errors.New("unexpected signing method")
	ErrInvalidTokenClaims      = errors.New("invalid token claims")
	ErrMissingHeaders          = errors.New("missing required headers")
	ErrNoHmacToken             = errors.New("hmac token is not configured")
	ErrTimestampParse          = errors.New("cannot parse timestamp")
	ErrTimestampExpired        = errors.New("timestamp expired or not yet valid")
	ErrInvalidSignature        = errors.New("invalid signature")
)
