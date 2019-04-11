package main

import (
	"database/sql"
	"log"

	errs "github.com/excelWithBusiness/MSGO-APIKEY-BE/errors"
	"github.com/excelWithBusiness/MSGO-APIKEY-BE/types"
)

// Services structure is used to contain the available rpc methods for time spent
type Services struct {
	db     *sql.DB
	helper *Helper
}

// SaveResult
func (m *Services) SaveResult(secretAuth string, clientConfig *types.APIKeyConfig) error {
	result, err := m.helper.ValidateAPIKey(m.db, secretAuth)
	if err != nil {
		log.Println(err)
		return errs.ErrInvalidAPIKey
	}

	*clientConfig = *result
	return nil
}

// GetClientConfig returns the json_config for an organisation (all orgs must share same config in the DB)
func (m *Services) GetResult(organisationID string, clientConfig *types.APIKeyConfig) error {
	result, err := m.helper.GetClientConfig(m.db, organisationID)
	if err != nil {
		log.Println(err)
		return errs.ErrInvalidAPIKey
	}

	*clientConfig = *result
	return nil
}
