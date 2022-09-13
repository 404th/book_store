package config

import "time"

const (
	// Database time layout
	DatabaseTimeLayout string = time.RFC3339
	// Access token expires in time
	AccessTokenExpiresInTime time.Duration = 1 * 24 * 60 * time.Minute
	// Refresh token expires in time
	RefreshTokenExpiresInTime time.Duration = 30 * 24 * 60 * time.Minute
)
