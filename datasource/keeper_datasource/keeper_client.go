package keeper_datasource

import (
	"encoding/base64"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/keeper-security/secrets-manager-go/core"
	ksm "github.com/keeper-security/secrets-manager-go/core"
)

// Constants for environment variables and Keeper record types.
// Unfortunately, we have to hardcode these values as Keeper doesn't provide a way to get them programmatically.
const (
	KEEPER_CONFIG_FILE_ENV_KEY  = "KEEPER_CONFIG_FILE"
	KEEPER_CONFIG_ENV_KEY       = "KSM_CONFIG"
	LOGIN_FIELD_TYPE            = "login"
	DATABASE_FIELD_TYPE         = "databaseCredentials"
	SERVER_FIELD_TYPE           = "serverCredentials"
	API_KEY_FIELD_TYPE          = "API Key"
	ENCRYPTED_NOTE_FIELD_TYPE   = "encryptedNotes"
	FILE_FIELD_TYPE             = "file"
	SOFTWARE_LICENSE_FIELD_TYPE = "softwareLicense"
	SSH_KEY_FIELD_TYPE          = "sshKeys"
	PASSWORD_FIELD_TYPE         = "password"
)

// Errors for handling configuration and record type issues.
var (
	ErrNoConfig        = errors.New("no config file specified in environment variable " + KEEPER_CONFIG_ENV_KEY + " and no config file set at " + KEEPER_CONFIG_FILE_ENV_KEY + " please set one of them")
	ErrWrongRecordType = errors.New("record is wrong type")
)

// KSMClient implemnents the KeeperClient interface and wraps the Keeper Secrets Manager client.
type KSMClient struct {
	KeeperClient *ksm.SecretsManager
}

// Interface for a Keeper client
type KeeperClient interface {
	GetSecret(uid string) (*ksm.Record, error)
	GetServerCredentials(r *ksm.Record) (*KeeperServerCredentials, error)
	GetDatabaseCredentials(r *ksm.Record) (*KeeperDataBaseCredentials, error)
	// Renamed for consistency: GetApiKey -> GetAPIKey to match KeeperAPIKey return type
	GetAPIKey(r *ksm.Record) (*KeeperAPIKey, error)
	GetEncryptedNote(r *ksm.Record) (*KeeperEncryptedNote, error)
	GetFile(r *ksm.Record) (*KeeperFile, error)
	GetSoftwareLicense(r *ksm.Record) (*KeeperSoftwareLicense, error)
	GetLogin(r *ksm.Record) (*KeeperLogin, error)
	GetSSHKey(r *ksm.Record) (*KeeperSSHKey, error)
}

// Convert KSMClient to KeeperClient interface (compile time check)
var _ KeeperClient = (*KSMClient)(nil)

// Returns a new Keeper client instance
func NewKeeperSecretClient() (*KSMClient, error) {
	// Get the client options based on environment variables.
	// Packer doesn't provide a way to pass the config via HCL2.
	clientOptions, err := getClientOptions()
	if err != nil {
		return nil, err
	}

	// Keeper doesn't validate the config content so we need to do it ourselves
	// without validation, the client will panic.
	if err := validateClientOptions(clientOptions); err != nil {
		return nil, err
	}

	// Create a new Keeper client with the provided options.
	ksmClient := ksm.NewSecretsManager(clientOptions)
	packerClient := &KSMClient{KeeperClient: ksmClient}

	return packerClient, nil
}

// GetServerCredentials retrieves a ServerCredentials record from Keeper
func (k *KSMClient) GetServerCredentials(r *ksm.Record) (*KeeperServerCredentials, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, SERVER_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the server credentials from the record
	return &KeeperServerCredentials{
		KeeperRecordField: *getRecordFields(record),
		HostConnection:    *getHostItemData(record),
		Login:             record.GetFieldValueByType("login"),
		Password:          record.GetFieldValueByType("password"),
	}, nil
}

// GetDatabaseCredentials retrieves a DatabaseCredentials record from Keeper
func (k *KSMClient) GetDatabaseCredentials(r *ksm.Record) (*KeeperDataBaseCredentials, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, DATABASE_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the database credentials from the record
	return &KeeperDataBaseCredentials{
		KeeperRecordField: *getRecordFields(record),
		HostConnection:    *getHostItemData(record),
		Login:             record.GetFieldValueByType("login"),
		Password:          record.GetFieldValueByType("password"),
		DbType:            record.GetFieldValueByType("text"),
	}, nil
}

// GetAPIKey retrieves an API Key record from Keeper
// Renamed from GetApiKey for consistency with KeeperAPIKey return type
func (k *KSMClient) GetAPIKey(r *ksm.Record) (*KeeperAPIKey, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, API_KEY_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the API key from the record
	return &KeeperAPIKey{
		KeeperRecordField: *getRecordFields(record),
		AppId:             record.GetFieldValueByLabel("AppID"),
		ClientSecret:      record.GetFieldValueByLabel("ClientSecret"),
	}, nil
}

// GetEncryptedNote retrieves an EncryptedNote record from Keeper
func (k *KSMClient) GetEncryptedNote(r *ksm.Record) (*KeeperEncryptedNote, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, ENCRYPTED_NOTE_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the encrypted note from the record
	noteContent := record.GetFieldValueByType("note")
	return &KeeperEncryptedNote{
		KeeperRecordField: *getRecordFields(record),
		Note:              noteContent,
		Date:              ConvertDateStr(record.GetFieldValueByType("date")),
	}, nil
}

// GetFile retrieves a File record from Keeper
func (k *KSMClient) GetFile(r *ksm.Record) (*KeeperFile, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, FILE_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the file record from the record
	return &KeeperFile{
		KeeperRecordField: *getRecordFields(record),
	}, nil
}

// GetSoftwareLicense retrieves a SoftwareLicense record from Keeper
func (k *KSMClient) GetSoftwareLicense(r *ksm.Record) (*KeeperSoftwareLicense, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, SOFTWARE_LICENSE_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the software license from the record
	return &KeeperSoftwareLicense{
		KeeperRecordField: *getRecordFields(record),
		LicenseNumber:     record.GetFieldValueByType("licenseNumber"),
		ActivationDate:    ConvertDateStr(record.GetFieldValueByType("date")),
		ExpirationDate:    ConvertDateStr(record.GetFieldValueByType("expirationDate")),
	}, nil
}

// GetLogin retrieves a Login record from Keeper
func (k *KSMClient) GetLogin(r *ksm.Record) (*KeeperLogin, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, LOGIN_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the login record from the record
	return &KeeperLogin{
		KeeperRecordField: *getRecordFields(record),
		Login:             record.GetFieldValueByType(LOGIN_FIELD_TYPE),
		Password:          record.GetFieldValueByType(PASSWORD_FIELD_TYPE),
		Url:               record.GetFieldValueByType("url"),
	}, nil

}

// GetSSHKey retrieves an SSH Key record from Keeper
func (k *KSMClient) GetSSHKey(r *ksm.Record) (*KeeperSSHKey, error) {
	// Validate the record is of the correct type
	record, err := k.validateRecord(r, SSH_KEY_FIELD_TYPE)
	if err != nil {
		return nil, err
	}

	// Extract the SSH key from the record
	return &KeeperSSHKey{
		KeeperRecordField: *getRecordFields(record),
		Login:             record.GetFieldValueByType(LOGIN_FIELD_TYPE),
		Passphrase:        record.GetFieldValueByType(PASSWORD_FIELD_TYPE),
		KeyPair:           *getKeyPairItemData(record),
		HostConnection:    *getHostItemData(record),
	}, nil
}

// GetSecret retrieves a generic record by uid from Keeper
func (k *KSMClient) GetSecret(uid string) (*ksm.Record, error) {
	// Fetch the record from Keeper using the provided uid
	records, err := k.KeeperClient.GetSecrets([]string{uid})
	if err != nil {
		return nil, err
	}

	// Check if any records were found, if not return an error
	if len(records) == 0 {
		return nil, fmt.Errorf("no records found for uid %s", uid)
	}

	// Get the first record from the list and return it (there should only be one)
	record := records[0]
	return record, nil
}

// ConvertDateStr converts a date (unix timestamp) to a time string
// the string field is misleading as it is actually a unix timestamp in milliseconds,
// but keeper coerces it to a string when returning the record value.
func ConvertDateStr(dateStr string) string {
	// Convert the date string to an int
	dateInt, err := strconv.Atoi(dateStr)
	if err != nil {
		return ""
	}

	// Convert the date int to a time.Time object and format it as a string
	timeStr := time.Unix(int64(dateInt/1000), 0).Format(time.RFC3339)
	return timeStr
}

// validateRecord checks if the record is of the expected type
func (k *KSMClient) validateRecord(r *ksm.Record, recordType string) (*ksm.Record, error) {
	if r.Type() != recordType {
		return nil, fmt.Errorf("%w Uid: %s ExpectedType: %s, ActualType: %s", ErrWrongRecordType, r.Uid, recordType, r.Type())
	}

	return r, nil
}

// getRecordFields extracts the common fields from a Keeper record
func getRecordFields(r *ksm.Record) *KeeperRecordField {
	return &KeeperRecordField{
		Uid:      r.Uid,
		Type:     r.Type(),
		Title:    r.Title(),
		Notes:    r.Notes(),
		FileRefs: getFileRecords(r),
	}
}

// getClientOptions retrieves the Keeper client options from environment variables
func getClientOptions() (*ksm.ClientOptions, error) {
	// Check if the KSM_CONFIG_FILE environment variable is set if so, read the file from disk
	// and initialize the client options with the file content.
	configFile, ok := os.LookupEnv(KEEPER_CONFIG_FILE_ENV_KEY)
	if ok {
		content, err := os.ReadFile(configFile)
		if err != nil {
			return nil, err
		}

		return &ksm.ClientOptions{
			Config: ksm.NewMemoryKeyValueStorage(string(content)),
		}, nil
	}

	// Check if the KSM_CONFIG environment variable is set, if so, use it to initialize the client options.
	// the environment variable content should be the base64 encoded config content
	configContent, ok := os.LookupEnv(KEEPER_CONFIG_ENV_KEY)
	if ok {
		return &ksm.ClientOptions{
			Config: ksm.NewMemoryKeyValueStorage(configContent),
		}, nil
	}

	return nil, ErrNoConfig
}

// validateClientOptions checks if the provided client options are valid
// keeper doesn't validate the config content so we need to do it ourselves
func validateClientOptions(c *ksm.ClientOptions) error {
	if c.Config.Get(core.KEY_APP_KEY) == "" || c.Config.Get(core.KEY_CLIENT_ID) == "" || c.Config.Get(core.KEY_PRIVATE_KEY) == "" {
		return fmt.Errorf("Invalid credentials - please provide a valid base64 encoded KSM config. One-time tokens are not allowed.")
	}

	return nil
}

// getFileRecords extracts the file records from a Keeper record
// this is a common record type that are part of all Keeper records.
func getFileRecords(r *ksm.Record) []FileRef {
	fileRefs := []FileRef{}
	for _, f := range r.Files {
		fileRef := FileRef{
			Uid:          f.Uid,
			Title:        f.Title,
			Name:         f.Name,
			Type:         f.Type,
			Size:         f.Size,
			LastModified: f.LastModified,
			Base64Data:   base64.StdEncoding.EncodeToString(f.GetFileData()),
		}

		fileRefs = append(fileRefs, fileRef)
	}

	return fileRefs
}

// getHostItemData extracts the host connection data from a Keeper record
func getHostItemData(secret *core.Record) *HostConnection {
	// Host data is stored in the host key in Keeper
	fields := secret.GetFieldsByType("host")
	if len(fields) == 0 {
		return &HostConnection{}
	}

	hc := &HostConnection{}

	// Make sure the field is a list of interface
	values, ok := fields[0]["value"].([]interface{})
	if !ok || len(values) == 0 {
		return hc
	}

	// Make sure the value is a map
	valuesMap, ok := values[0].(map[string]interface{})
	if !ok {
		return hc
	}

	// Extract the host connection data from the map
	if val, ok := valuesMap["hostName"].(string); ok {
		hc.HostName = val
	}

	// Extract the port from the map
	if val, ok := valuesMap["port"].(string); ok {
		// Attempt to convert the port from a string to an int
		// if it fails return -1
		port, err := strconv.Atoi(val)
		if err != nil {
			hc.Port = -1
			return hc
		}

		hc.Port = port
	}

	return hc
}

// getKeyPairItemData extracts the key pair data from a Keeper SSH key
func getKeyPairItemData(secret *core.Record) *KeyPair {
	// Key pair data is stored in the keyPair key in Keeper
	fields := secret.GetFieldsByType("keyPair")
	if len(fields) == 0 {
		return &KeyPair{}
	}

	keypair := &KeyPair{}

	// Make sure the field is a list of interface
	values, ok := fields[0]["value"].([]interface{})
	if !ok || len(values) == 0 {
		return keypair
	}

	// Make sure the value is a map
	valuesMap, ok := values[0].(map[string]interface{})
	if !ok {
		return keypair
	}

	// Extract the key pair data from the map
	if pub, ok := valuesMap["publicKey"].(string); ok {
		keypair.PublicKey = pub
	}

	if priv, ok := valuesMap["privateKey"].(string); ok {
		keypair.PrivateKey = priv
	}

	return keypair
}
