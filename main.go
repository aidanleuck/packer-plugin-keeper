// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package main

import (
	"fmt"
	"os"

	keeper_api_key "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-api-key"
	keeper_database_credentials "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-database-credentials"
	keeper_encrypted_note "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-encrypted-note"
	keeper_file "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-file"
	keeper_login "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-login"
	keeper_server_credentials "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-server-credentials"
	keeper_software_license "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-software-license"
	keeper_ssh_key "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-ssh-key"
	version "github.com/aidanleuck/packer-plugin-keeper/version"

	"github.com/hashicorp/packer-plugin-sdk/plugin"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterDatasource("login", new(keeper_login.Datasource))
	pps.RegisterDatasource("software-license", new(keeper_software_license.Datasource))
	pps.RegisterDatasource("encrypted-note", new(keeper_encrypted_note.Datasource))
	pps.RegisterDatasource("file", new(keeper_file.Datasource))
	pps.RegisterDatasource("ssh-key", new(keeper_ssh_key.Datasource))
	pps.RegisterDatasource("api-key", new(keeper_api_key.Datasource))
	pps.RegisterDatasource("database-credential", new(keeper_database_credentials.Datasource))
	pps.RegisterDatasource("server-credential", new(keeper_server_credentials.Datasource))

	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
