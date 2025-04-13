// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_database_credentials

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v datasource/keeper_datasource/keeper-database-credentials/data_keeper_database_credentials_acc_test.go -timeout=120m
// TestAccKeeperDatabaseCredentials is an acceptance test for the Keeper Database Credentials datasource. It tests against a real Keeper vault and you must have the ability
// to contact the Keeper API. It is recommended to use a test vault for this test.
func TestAccKeeperDatabaseCredentials(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_database_credential_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "keeper-database",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			logLines := []string{
				"null.basic-example: Title: test-db",
				"null.basic-example: Notes: coolio-db",
				"null.basic-example: Login: test-user",
				"null.basic-example: Password: test-password",
				"null.basic-example: Host: 127.0.0.1",
				"null.basic-example: Port: 5432",
				"null.basic-example: Type: postgres",
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
