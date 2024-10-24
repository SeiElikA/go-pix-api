package config

import (
	"net/http"
	"strings"
)

type Config struct {
	ServerURL string
}

var AppConfig *Config

func InitializeConfig(r *http.Request) {
	AppConfig = &Config{
		ServerURL: getServerURL(r),
	}
}

func getServerURL(r *http.Request) string {
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}

	host := r.Host

	return strings.Join([]string{scheme, "://", host, "/"}, "")
}
