// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_api_key

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/keeper_datasource/keeper-api-key/data_keeper_api_key_acc_test.go  -timeout=120m
// This is an integration test that pulls a real secret from Keeper and checks the output. Don't use real secrets in this test.
func TestAccKeeperApiKey(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_login_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "keeper-api-key",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			logLines := []string{
				"null.basic-example: Title: Test API Key",
				"null.basic-example: ClientSecret: 12345",
				"null.basic-example: AppID: test-app-id",
				"null.basic-example: Notes: Api key",	
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
