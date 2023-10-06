package config

var defaultConfig = map[string]interface{}{
	"auth.refresh_subject":        RefreshTokenSubject,
	"auth.access_subject":         AccessTokenSubject,
	"auth.access_expirationTime":  AccessTokenExpireDuration,
	"auth.refresh_expirationTime": RefreshTokenExpireDuration,
}
