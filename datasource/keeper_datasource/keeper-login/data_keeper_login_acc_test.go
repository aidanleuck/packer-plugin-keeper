// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_login

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/scaffolding/data_acc_test.go  -timeout=120m
// TestAccKeeperLogin is an integration test that pulls a Keeper login from a real Keeper account and checks the output. Don't use real secrets in this test.
func TestAccKeeperLogin(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_login_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "keeper-login",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			logLines := []string{
				"null.basic-example: Login: test@selinc.com",
				"null.basic-example: Password: testing123",
				"null.basic-example: Url: https://test.com",
				"null.basic-example: Notes: Hello",	
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
