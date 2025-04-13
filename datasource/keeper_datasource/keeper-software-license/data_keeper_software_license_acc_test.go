// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_software_license

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/keeper_datasource/keeper-software-license/data_keeper_software_license_acc_test.go  -timeout=120m
// TestAccKeeperSoftwareLicense is an integration test that pulls a real secret from Keeper and checks the output. Don't use real secrets in this test.
func TestAccKeeperSoftwareLicense(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_server_credential_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "keeper-software-license",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			logLines := []string{
				"null.basic-example: Title: Test License",
				"null.basic-example: License Key: 12345",
				"null.basic-example: Expiration Date: 2025-01-01T00:00:00-07:00",
				"null.basic-example: Activation Date: 2025-01-02T00:00:00-07:00",
				"null.basic-example: Notes: best license",
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
