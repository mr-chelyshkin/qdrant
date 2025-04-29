package auth

import (
	"crypto/hmac"
	"encoding/hex"
	"strconv"
	"time"

	sha256 "github.com/minio/sha256-simd"
	"github.com/pkg/errors"
)

// HMACAuth ...
type HMACAuth struct {
	secret []byte
}

// NewHMACAuth ...
func NewHMACAuth(secret string) *HMACAuth {
	return &HMACAuth{
		secret: []byte(secret),
	}
}

// GenerateSignature ...
func (h *HMACAuth) GenerateSignature(data string) string {
	mac := hmac.New(sha256.New, h.secret)
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// ValidateRequest ...
func (h *HMACAuth) ValidateRequest(timestamp, signature string) error {
	if timestamp == "" || signature == "" {
		return ErrMissingHeaders
	}
	if len(h.secret) == 0 {
		return ErrNoHmacToken
	}

	ts, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		return errors.Wrap(ErrTimestampParse, err.Error())
	}
	currentTime := time.Now().UTC().Unix()
	if currentTime-ts > TimestampDelay || ts > currentTime+TimestampDelay {
		return ErrTimestampExpired
	}

	expectedSignature := h.GenerateSignature(timestamp)
	if !hmac.Equal([]byte(signature), []byte(expectedSignature)) {
		return ErrInvalidSignature
	}
	return nil
}
