// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package keeper_file

import (
	_ "embed"
	"encoding/base64"
	"fmt"
	"os/exec"
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/packer-plugin-sdk/acctest"
)

//go:embed test-fixtures/template.pkr.hcl
var testDatasourceHCL2Basic string

// Run with: PACKER_ACC=1 go test -count 1 -v ./datasource/keeper_datasource/keeper-file/data_keeper_file_acc_test.go  -timeout=120m
// TestAccKeeperFile is an integration test that pulls a real secret from Keeper and checks the output. Don't use real secrets in this test.
func TestAccKeeperFile(t *testing.T) {
	testCase := &acctest.PluginTestCase{
		Name: "keeper_login_basic_test",
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
				"null.basic-example: Title: Single File",
				"null.basic-example: Size: 5",
				"null.basic-example: Last Modified: 1680294505324",
				"null.basic-example: Notes: Single test",
				"null.basic-example: File Name: test1.txt",
				"null.basic-example: Type: text/plain",
				"null.basic-example: File UID: mT69r6vfdC7lg7ERvpc0lQ",	
				fmt.Sprintf("null.basic-example: Base 64 Data: %s", base64.StdEncoding.EncodeToString([]byte("test1"))),
			}

			if err := keeper_datasource.RunPackerAcceptanceTest(t, buildCommand, logfile, logLines); err != nil {
				return err
			}

			return nil
		},
	}
	acctest.TestPlugin(t, testCase)
}
