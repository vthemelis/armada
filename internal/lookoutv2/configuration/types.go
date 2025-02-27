package configuration

import (
	"time"

	"github.com/armadaproject/armada/internal/armada/configuration"
)

type LookoutV2Configuration struct {
	ApiPort            int
	CorsAllowedOrigins []string
	Tls                TlsConfig

	Postgres configuration.PostgresConfig

	PrunerConfig PrunerConfig
}

type TlsConfig struct {
	Enabled  bool
	KeyPath  string
	CertPath string
}

type PrunerConfig struct {
	ExpireAfter time.Duration
	Timeout     time.Duration
	BatchSize   int
}
