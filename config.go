package xingyun

import (
	"time"
)

const (
	DefaultSecret = "D893ACDB5B524C6X"
)

type Config struct {
	EnableDebug bool

	CookieSecret   string
	SessionTimeout int32

	StaticDir       string
	StaticPrefix    string
	StaticIndexFile string

	StaticHost            string
	StaticHostExcludeType string
	StaticHostExcludeFile string

	XSRFSecret         string
	XSRFAllowedOrigins []string
}

func setDefaultConfig(config *Config) {
	if config.StaticIndexFile == "" {
		config.StaticIndexFile = "index.html"
	}
	if config.StaticDir == "" {
		config.StaticDir = "static"
	}
	if config.CookieSecret == "" {
		config.CookieSecret = DefaultSecret
	}
	if config.XSRFSecret == "" {
		config.XSRFSecret = DefaultSecret
	}
	if config.SessionTimeout == 0 {
		// 7 days
		config.SessionTimeout = 7 * 24 * int32(time.Hour.Minutes())
	}
}
