package transport

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestStreamableHTTP_WithOAuth(t *testing.T) {
	// Create a test server that requires OAuth
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader != "Bearer test-token" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Return a successful response
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]any{
			"jsonrpc": "2.0",
			"id":      1,
			"result":  "success",
		})
	}))
	defer server.Close()

	// Create a token store with a valid token
	tokenStore := NewMemoryTokenStore()
	validToken := &Token{
		AccessToken:  "test-token",
		TokenType:    "Bearer",
		RefreshToken: "refresh-token",
		ExpiresIn:    3600,
		ExpiresAt:    time.Now().Add(1 * time.Hour), // Valid for 1 hour
	}
	tokenStore.SaveToken(validToken)

	// Create OAuth config
	oauthConfig := OAuthConfig{
		ClientID:     "test-client",
		RedirectURI:  "http://localhost:8085/callback",
		Scopes:       []string{"mcp.read", "mcp.write"},
		TokenStore:   tokenStore,
		PKCEEnabled:  true,
	}

	// Create StreamableHTTP with OAuth
	transport, err := NewStreamableHTTP(server.URL, WithOAuth(oauthConfig))
	if err != nil {
		t.Fatalf("Failed to create StreamableHTTP: %v", err)
	}

	// Verify that OAuth is enabled
	if !transport.IsOAuthEnabled() {
		t.Errorf("Expected IsOAuthEnabled() to return true")
	}
	
	// Verify the OAuth handler is set
	if transport.GetOAuthHandler() == nil {
		t.Errorf("Expected GetOAuthHandler() to return a handler")
	}
}

func TestStreamableHTTP_WithOAuth_Unauthorized(t *testing.T) {
	// Create a test server that requires OAuth
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Always return unauthorized
		w.WriteHeader(http.StatusUnauthorized)
	}))
	defer server.Close()

	// Create an empty token store
	tokenStore := NewMemoryTokenStore()

	// Create OAuth config
	oauthConfig := OAuthConfig{
		ClientID:     "test-client",
		RedirectURI:  "http://localhost:8085/callback",
		Scopes:       []string{"mcp.read", "mcp.write"},
		TokenStore:   tokenStore,
		PKCEEnabled:  true,
	}

	// Create StreamableHTTP with OAuth
	transport, err := NewStreamableHTTP(server.URL, WithOAuth(oauthConfig))
	if err != nil {
		t.Fatalf("Failed to create StreamableHTTP: %v", err)
	}

	// Send a request
	_, err = transport.SendRequest(context.Background(), JSONRPCRequest{
		JSONRPC: "2.0",
		ID:      1,
		Method:  "test",
	})

	// Verify the error is an OAuthAuthorizationRequiredError
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	oauthErr, ok := err.(*OAuthAuthorizationRequiredError)
	if !ok {
		t.Fatalf("Expected OAuthAuthorizationRequiredError, got %T: %v", err, err)
	}

	// Verify the error has the handler
	if oauthErr.Handler == nil {
		t.Errorf("Expected OAuthAuthorizationRequiredError to have a handler")
	}
}

func TestStreamableHTTP_IsOAuthEnabled(t *testing.T) {
	// Create StreamableHTTP without OAuth
	transport1, err := NewStreamableHTTP("http://example.com")
	if err != nil {
		t.Fatalf("Failed to create StreamableHTTP: %v", err)
	}

	// Verify OAuth is not enabled
	if transport1.IsOAuthEnabled() {
		t.Errorf("Expected IsOAuthEnabled() to return false")
	}

	// Create StreamableHTTP with OAuth
	transport2, err := NewStreamableHTTP("http://example.com", WithOAuth(OAuthConfig{
		ClientID: "test-client",
	}))
	if err != nil {
		t.Fatalf("Failed to create StreamableHTTP: %v", err)
	}

	// Verify OAuth is enabled
	if !transport2.IsOAuthEnabled() {
		t.Errorf("Expected IsOAuthEnabled() to return true")
	}
}