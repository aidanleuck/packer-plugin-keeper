package keeper_datasource

import (
	"sync"
)

var (
	globalPackerSecretsManager *PackerKeeperClient
	once                       sync.Once
)

// PackerKeeperClient is a wrapper around the KeeperClient interface
type PackerKeeperClient struct {
	KeeperClient KeeperClient
}

// NewClient creates a new PackerKeeperClient
func NewClient(c KeeperClient) *PackerKeeperClient {
	return &PackerKeeperClient{
		KeeperClient: c,
	}
}

// GetSecretClient is a singleton Keeper client that is shared
// across datasources. This is to prevent each datasource call in
// a Packer build to reinitialize the client.
func GetSecretClient() (*PackerKeeperClient, error) {
	// error to return if initialization fails
	var initError error
	once.Do(func() {
		// We initialize a real client here, but we could also modify this function
		// to pass in a mock client for testing purposes.
		sc, err := NewKeeperSecretClient()
		if err != nil {
			initError = err
			return
		}

		// Set the global variable
		globalPackerSecretsManager = &PackerKeeperClient{KeeperClient: sc}
	})

	// If there were any errors initializing the client return nil and fail.
	if initError != nil {
		return nil, initError
	}

	return globalPackerSecretsManager, nil
}

// GetServerCredentials retrieves the server credentials for a given uid
func (c *PackerKeeperClient) GetServerCredentials(uid string) (*KeeperServerCredentials, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}

	return c.KeeperClient.GetServerCredentials(r)
}

// GetDatabaseCredentials retrieves the database credentials for a given uid
func (c *PackerKeeperClient) GetDatabaseCredentials(uid string) (*KeeperDataBaseCredentials, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetDatabaseCredentials(r)
}

// GetAPIKey retrieves the API key for a given uid
func (c *PackerKeeperClient) GetAPIKey(uid string) (*KeeperAPIKey, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetAPIKey(r)
}

// GetEncryptedNote retrieves the encrypted note for a given uid
func (c *PackerKeeperClient) GetEncryptedNote(uid string) (*KeeperEncryptedNote, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetEncryptedNote(r)
}

// GetFile retrieves the file for a given uid
func (c *PackerKeeperClient) GetFile(uid string) (*KeeperFile, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetFile(r)
}

// GetSoftwareLicense retrieves the software license for a given uid
func (c *PackerKeeperClient) GetSoftwareLicense(uid string) (*KeeperSoftwareLicense, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetSoftwareLicense(r)
}

// GetLogin retrieves the login for a given uid
func (c *PackerKeeperClient) GetLogin(uid string) (*KeeperLogin, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetLogin(r)
}

// GetSSHKey retrieves the SSH key for a given uid
func (c *PackerKeeperClient) GetSSHKey(uid string) (*KeeperSSHKey, error) {
	r, err := c.KeeperClient.GetSecret(uid)
	if err != nil {
		return nil, err
	}
	return c.KeeperClient.GetSSHKey(r)
}
