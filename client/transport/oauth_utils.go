package transport

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

// GenerateRandomString generates a random string of the specified length
func GenerateRandomString(length int) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes)[:length], nil
}

// GenerateCodeVerifier generates a code verifier for PKCE
func GenerateCodeVerifier() (string, error) {
	// According to RFC 7636, the code verifier should be between 43 and 128 characters
	return GenerateRandomString(64)
}

// GenerateCodeChallenge generates a code challenge from a code verifier
func GenerateCodeChallenge(codeVerifier string) string {
	// SHA256 hash the code verifier
	hash := sha256.Sum256([]byte(codeVerifier))
	// Base64url encode the hash
	return base64.RawURLEncoding.EncodeToString(hash[:])
}

// GenerateState generates a state parameter for OAuth
func GenerateState() (string, error) {
	return GenerateRandomString(32)
}

// ValidateRedirectURI validates that a redirect URI is secure
func ValidateRedirectURI(redirectURI string) error {
	// According to the spec, redirect URIs must be either localhost URLs or HTTPS URLs
	if redirectURI == "" {
		return fmt.Errorf("redirect URI cannot be empty")
	}

	// Check if it's a localhost URL
	if len(redirectURI) >= 9 && redirectURI[:9] == "http://lo" {
		return nil
	}

	// Check if it's an HTTPS URL
	if len(redirectURI) >= 8 && redirectURI[:8] == "https://" {
		return nil
	}

	return fmt.Errorf("redirect URI must be either a localhost URL or an HTTPS URL")
}