// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_server_credentials

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/keeper_datasource/keeper-file/data_keeper_file_acc_test.go  -timeout=120m
// TestAccKeeperServerCredentials is an integration test that pulls a real secret from Keeper and checks the output. Don't use real secrets in this test.
func TestAccKeeperServerCredentials(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_server_credential_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "keeper-server-credential",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			logLines := []string{
				"null.basic-example: Title: test-server-credential",
				"null.basic-example: Notes: my-server",
				"null.basic-example: Login: test-user",
				"null.basic-example: Password: test-password",
				"null.basic-example: Host: 127.0.0.1",
				"null.basic-example: Port: 443",
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
