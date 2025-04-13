// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_encrypted_note

import (
	_ "embed"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/keeper_datasource/keeper-encrypted-note/data_keeper_encrypted_note_acc_test.go  -timeout=120m
// TestAccKeeperEncryptedNote is an integration test that pulls a real secret from Keeper and checks the output. Don't use real secrets in this test.
func TestAccKeeperEncryptedNote(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_encrypted_note_basic_test",
		Setup: func() error {
			return nil
		},
		Teardown: func() error {
			return nil
		},
		Template: testDatasourceHCL2Basic,
		Type:     "keeper-file",
		Check: func(buildCommand *exec.Cmd, logfile string) error {
			logLines := []string{
				"null.basic-example: Title: test-note",
				"null.basic-example: notes: My test note",
				"null.basic-example: securedNote: super-secret-string",
				"null.basic-example: date: 1999-08-12T00:00:00-06:00",
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
