package main

import (
	"database/sql"
	"github.com/excelWithBusiness/MSGO-APIKEY-BE/types"
)

// Helper structure to map database into structure types
type Helper int

// ValidateAPIKey pull aorganisation_id, group_id, api_key, api_secret by APIKEY
func (h *Helper) ValidateAPIKey(db *sql.DB, APIKey string) (*types.APIKeyConfig, error) {
	result := types.APIKeyConfig{}

	return &result, nil
}

// GetClientConfig pull aorganisation_id, group_id, api_key, api_secret by APIKEY
func (h *Helper) GetClientConfig(db *sql.DB, organisationID string) (*types.APIKeyConfig, error) {
	result := types.APIKeyConfig{}

	return &result, nil
}
