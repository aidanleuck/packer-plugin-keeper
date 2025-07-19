package common

import (
	"testing"

	"github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource"
	keeper_api_key "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-api-key"
	keeper_database_credentials "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-database-credentials"
	keeper_encrypted_note "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-encrypted-note"
	keeper_file "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-file"
	keeper_login "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-login"
	keeper_server_credentials "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-server-credentials"
	keeper_software_license "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-software-license"
	keeper_ssh_key "github.com/aidanleuck/packer-plugin-keeper/datasource/keeper_datasource/keeper-ssh-key"
	"github.com/hashicorp/packer-plugin-sdk/packer"
	"github.com/stretchr/testify/require"
)

type tc struct {
	TestName   string
	DataSource packer.Datasource
}

// TestInvalidConfigReturnsError tests that the Configure method returns an error when the config is invalid.
// in this case it is invalid because the UID is not set.
func TestInvalidConfigReturnsError(t *testing.T) {
	tcs := []tc{
		{
			DataSource: &keeper_api_key.Datasource{},
			TestName:   "keeper_api_key",
		},
		{
			DataSource: &keeper_database_credentials.Datasource{},
			TestName:   "keeper_database_credentials",
		},
		{
			DataSource: &keeper_encrypted_note.Datasource{},
			TestName:   "keeper_encrypted_note",
		},
		{
			DataSource: &keeper_file.Datasource{},
			TestName:   "keeper_file",
		},
		{
			DataSource: &keeper_login.Datasource{},
			TestName:   "keeper_login",
		},
		{
			DataSource: &keeper_server_credentials.Datasource{},
			TestName:   "keeper_server_credentials",
		},
		{
			DataSource: &keeper_software_license.Datasource{},
			TestName:   "keeper_software_license",
		},
		{
			DataSource: &keeper_ssh_key.Datasource{},
			TestName:   "keeper_ssh_key",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.TestName, func(t *testing.T) {
			err := tc.DataSource.Configure()
			require.ErrorIs(t, err, keeper_datasource.ErrUidRequired)
		})
	}
}

// TestValidConfigReturnsNoError tests that the Configure method returns no error when uid is set.
func TestValidConfigReturnsNoError(t *testing.T) {
	testUid := "test-uid"
	config := &keeper_datasource.Config{
		Uid: &testUid,
	}
	tcs := []tc{
		{
			DataSource: &keeper_api_key.Datasource{
				Config: *config,
			},
			TestName: "keeper_api_key",
		},
		{
			DataSource: &keeper_database_credentials.Datasource{
				Config: *config,
			},
			TestName: "keeper_database_credentials",
		},
		{
			DataSource: &keeper_encrypted_note.Datasource{
				Config: *config,
			},
			TestName: "keeper_encrypted_note",
		},
		{
			DataSource: &keeper_file.Datasource{
				Config: *config,
			},
			TestName: "keeper_file",
		},
		{
			DataSource: &keeper_login.Datasource{
				Config: *config,
			},
			TestName: "keeper_login",
		},
		{
			DataSource: &keeper_server_credentials.Datasource{
				Config: *config,
			},
			TestName: "keeper_server_credentials",
		},
		{
			DataSource: &keeper_software_license.Datasource{
				Config: *config,
			},
			TestName: "keeper_software_license",
		},
		{
			DataSource: &keeper_ssh_key.Datasource{
				Config: *config,
			},
			TestName: "keeper_ssh_key",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.TestName, func(t *testing.T) {
			err := tc.DataSource.Configure()
			require.NoError(t, err)
		})
	}
}
