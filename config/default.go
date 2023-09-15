package config

var defaultConfig = map[string]interface{}{
	"auth.refresh_subject": RefreshTokenSubject,
	"auth.accdess_subject": AccessTokenSubject,
}
