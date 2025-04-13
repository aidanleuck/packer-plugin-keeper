package keeper_datasource

import (
	"github.com/keeper-security/secrets-manager-go/core"
	ksm "github.com/keeper-security/secrets-manager-go/core"
	"github.com/stretchr/testify/mock"
)

// MockWrapper wraps a real KeeperClient and allows mocking the GetSecret method.
type MockKeeperClient struct {
    mock.Mock
    TestClient KeeperClient
}

var _ KeeperClient = (*MockKeeperClient)(nil)

// Mock only GetSecret, delegate the rest to the real client
func (m *MockKeeperClient) GetSecret(uid string) (*core.Record, error) {
    args := m.Called()
    return args.Get(0).(*core.Record), args.Error(1)
}

// Delegate the rest of the methods to the real client
func (m *MockKeeperClient) GetLogin(r *ksm.Record) (*KeeperLogin, error) {
    return m.TestClient.GetLogin(r)
}

func (m *MockKeeperClient) GetApiKey(r *ksm.Record) (*KeeperAPIKey, error) {
    return m.TestClient.GetApiKey(r)
}

func (m *MockKeeperClient) GetSoftwareLicense(r *ksm.Record) (*KeeperSoftwareLicense, error) {
    return m.TestClient.GetSoftwareLicense(r)
}

func (m *MockKeeperClient) GetFile(r *ksm.Record) (*KeeperFile, error) {
    return m.TestClient.GetFile(r)
}

func (m *MockKeeperClient) GetEncryptedNote(r *ksm.Record) (*KeeperEncryptedNote, error) {
    return m.TestClient.GetEncryptedNote(r)
}

func (m *MockKeeperClient) GetDatabaseCredentials(r *ksm.Record) (*KeeperDataBaseCredentials, error) {
    return m.TestClient.GetDatabaseCredentials(r)
}

func (m *MockKeeperClient) GetServerCredentials(r *ksm.Record) (*KeeperServerCredentials, error) {
    return m.TestClient.GetServerCredentials(r)
}

func (m *MockKeeperClient) GetSSHKey(r *ksm.Record) (*KeeperSSHKey, error) {
    return m.TestClient.GetSSHKey(r)
}