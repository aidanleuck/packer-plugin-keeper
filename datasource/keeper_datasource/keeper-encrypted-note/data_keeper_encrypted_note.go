// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:generate packer-sdc mapstructure-to-hcl2 -type DatasourceOutput
package keeper_encrypted_note

import (
	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	keeper "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/hashicorp/packer-plugin-sdk/hcl2helper"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/hashicorp/packer-plugin-sdk/template/config"
	"github.com/zclconf/go-cty/cty"
)

type Datasource struct {
	Config keeper_datasource.Config
}

type DatasourceOutput struct {
	keeper.KeeperEncryptedNote `mapstructure:",squash"`
}

// ConfigSpec converts the config struct to a spec for HCL2
func (d *Datasource) ConfigSpec() hcldec.ObjectSpec {
	return d.Config.FlatMapstructure().HCL2Spec()
}

// Configure decodes the raw configuration into the Datasource struct
func (d *Datasource) Configure(raws ...interface{}) error {
	err := config.Decode(&d.Config, nil, raws...)
	if err != nil {
		return err
	}

	// Validate all required fields are set and valid
	if err := keeper_datasource.ValidateDataSourceConfig(d.Config); err != nil {
		return err
	}

	return nil
}

// OutputSpec converts the output struct to a spec for HCL2
func (d *Datasource) OutputSpec() hcldec.ObjectSpec {
	return (&DatasourceOutput{}).FlatMapstructure().HCL2Spec()
}

// Execute fetches the encrypted note from Keeper and returns it as a cty.Value
func (d *Datasource) Execute() (cty.Value, error) {
	// Get the Keeper client
	keeperClient, err := keeper.GetSecretClient()
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	// Fetch the encrypted note using the UID from the config
	note, err := keeperClient.GetEncryptedNote(*d.Config.Uid)
	if err != nil {
		return cty.NullVal(cty.EmptyObject), err
	}

	// Set the log secret filter to mask sensitive information
	packersdk.LogSecretFilter.Set(note.Note)
	output := &DatasourceOutput{
		KeeperEncryptedNote: *note,
	}

	return hcl2helper.HCL2ValueFromConfig(output, d.OutputSpec()), nil
}
