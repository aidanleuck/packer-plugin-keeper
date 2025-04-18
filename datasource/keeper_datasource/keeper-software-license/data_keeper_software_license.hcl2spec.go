// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package keeper_software_license

import (
	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatDatasourceOutput struct {
	Uid            *string                         `mapstructure:"uid" cty:"uid" hcl:"uid"`
	Type           *string                         `mapstructure:"type" cty:"type" hcl:"type"`
	Title          *string                         `mapstructure:"title" cty:"title" hcl:"title"`
	Notes          *string                         `mapstructure:"notes" cty:"notes" hcl:"notes"`
	FileRefs       []keeper_datasource.FlatFileRef `mapstructure:"file_refs" cty:"file_refs" hcl:"file_refs"`
	LicenseNumber  *string                         `mapstructure:"license_number" cty:"license_number" hcl:"license_number"`
	ActivationDate *string                         `mapstructure:"activation_date" cty:"activation_date" hcl:"activation_date"`
	ExpirationDate *string                         `mapstructure:"expiration_date" cty:"expiration_date" hcl:"expiration_date"`
}

// FlatMapstructure returns a new FlatDatasourceOutput.
// FlatDatasourceOutput is an auto-generated flat version of DatasourceOutput.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*DatasourceOutput) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatDatasourceOutput)
}

// HCL2Spec returns the hcl spec of a DatasourceOutput.
// This spec is used by HCL to read the fields of DatasourceOutput.
// The decoded values from this spec will then be applied to a FlatDatasourceOutput.
func (*FlatDatasourceOutput) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"uid":             &hcldec.AttrSpec{Name: "uid", Type: cty.String, Required: false},
		"type":            &hcldec.AttrSpec{Name: "type", Type: cty.String, Required: false},
		"title":           &hcldec.AttrSpec{Name: "title", Type: cty.String, Required: false},
		"notes":           &hcldec.AttrSpec{Name: "notes", Type: cty.String, Required: false},
		"file_refs":       &hcldec.BlockListSpec{TypeName: "file_refs", Nested: hcldec.ObjectSpec((*keeper_datasource.FlatFileRef)(nil).HCL2Spec())},
		"license_number":  &hcldec.AttrSpec{Name: "license_number", Type: cty.String, Required: false},
		"activation_date": &hcldec.AttrSpec{Name: "activation_date", Type: cty.String, Required: false},
		"expiration_date": &hcldec.AttrSpec{Name: "expiration_date", Type: cty.String, Required: false},
	}
	return s
}
