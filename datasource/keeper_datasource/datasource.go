package keeper_datasource

import (
	"errors"
)

var (
	ErrUidRequired = errors.New("uid is a required field")
)

// ValidateDataSourceConfig validates the configuration for all Keeper datasources.
// uid is a required field for all datasources and must be set.
func ValidateDataSourceConfig(config Config) error {
	if config.Uid == nil || *config.Uid == "" {
		return ErrUidRequired
	}

	return nil
}
