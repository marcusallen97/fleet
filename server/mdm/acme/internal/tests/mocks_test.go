package tests

import (
	"context"
	"errors"

	"github.com/fleetdm/fleet/v4/server/fleet"
)

// Mock implementations for dependencies outside the bounded context

// mockDataProviders combines all provider interfaces for testing.
type mockDataProviders struct {
	appCfg *fleet.AppConfig
}

func newMockDataProviders(appCfg *fleet.AppConfig) *mockDataProviders {
	return &mockDataProviders{appCfg: appCfg}
}

func (m *mockDataProviders) AppConfig(ctx context.Context) (*fleet.AppConfig, error) {
	return m.appCfg, nil
}

// Returns a valid row with `valid-serial` else no row
func (m *mockDataProviders) GetHostDEPAssignmentsBySerial(ctx context.Context, serial string) ([]*fleet.HostDEPAssignment, error) {
	// For testing, we can return a fixed response or an empty slice based on the serial number
	if serial == "valid-serial" {
		return []*fleet.HostDEPAssignment{
			{
				HostID: 1,
			},
		}, nil
	} else if serial == "error-serial" {
		return nil, errors.New("Mocked error for GetHostDEPAssignmentsBySerial")
	}
	return []*fleet.HostDEPAssignment{}, nil
}
