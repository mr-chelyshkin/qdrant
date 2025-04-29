package auth

import "time"

const (
	TimestampDelay  = 15
	HeaderTimestamp = "x-timestamp"
	HeaderSignature = "x-signature"
	HeaderAuth      = "Authorization"
)

// Authenticator ...
type Authenticator struct {
	hmacToken string
	jwtSecret []byte
	hmac      *HMACAuth
	jwt       *JWTAuth
}

// NewAuthenticator ...
func NewAuthenticator(hmacToken string, jwtSecret string) *Authenticator {
	auth := &Authenticator{
		hmacToken: hmacToken,
		jwtSecret: []byte(jwtSecret),
	}
	auth.hmac = NewHMACAuth(hmacToken)
	auth.jwt = NewJWTAuth(jwtSecret)
	return auth
}

// ValidateHmacRequest ...
func (a *Authenticator) ValidateHmacRequest(timestamp, signature string) error {
	return a.hmac.ValidateRequest(timestamp, signature)
}

// ValidateJWTToken ...
func (a *Authenticator) ValidateJWTToken(tokenString string) (*Claims, error) {
	return a.jwt.ValidateToken(tokenString)
}

// GenerateToken ...
func (a *Authenticator) GenerateToken(userID int, role Role, expiresIn time.Duration) (string, error) {
	return a.jwt.GenerateToken(userID, role, expiresIn)
}
