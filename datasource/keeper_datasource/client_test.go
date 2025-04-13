package keeper_datasource

import (
	"bytes"
	"fmt"
	"html/template"
	"testing"
	"time"

	ksm "github.com/keeper-security/secrets-manager-go/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// TestInvalidConfigReturnsError tests that each datasource returns an error when the type of secret passed is not
// the same type as the datasource.
func TestInvalidTypeThrowsError(t *testing.T){
	// No type should have a invalid type so this test will work on all datasources.
	typeRecord := `{
		"type": "invalid"
		}
		`
	client := getMockedClient(typeRecord)

	type tc struct{
		TestName string
		function func(string) (interface{}, error)
	} 
	
	// Test each data source with a type of invalid. This should return an error for all datasources.
	tcs := []tc{
		{
			TestName: "GetLogin",
			function: func(uid string) (interface{}, error) {
				return client.GetLogin(uid)
			},
		},
		{
			TestName: "GetApiKey",
			function: func(uid string) (interface{}, error) {
				return client.GetApiKey(uid)
			},
		},
		{
			TestName: "GetSoftwareLicense",
			function: func(uid string) (interface{}, error) {
				return client.GetSoftwareLicense(uid)
			},
		},
		{
			TestName: "GetFile",
			function: func(uid string) (interface{}, error) {	
				return client.GetFile(uid)
			},
		},
		{
			TestName: "GetEncryptedNote",
			function: func(uid string) (interface{}, error) {		
				return client.GetEncryptedNote(uid)
			},
		},
		{	
			TestName: "GetDatabaseCredentials",
			function: func(uid string) (interface{}, error) {
				return client.GetDatabaseCredentials(uid)
			},
		},
		{
			TestName: "GetServerCredentials",
			function: func(uid string) (interface{}, error) {
				return client.GetServerCredentials(uid)
			},
		},
		{	
			TestName: "GetSSHKey",
			function: func(uid string) (interface{}, error) {
				return client.GetSSHKey(uid)
			},
		},	
	}

	// Assert that wrong record type error is returned for each data source.
	for _, tc := range tcs {
		t.Run(tc.TestName, func(t *testing.T) {
			_, err := tc.function("test-uid")
			assert.ErrorIs(t, err, ErrWrongRecordType)
		})
	}
}

// TestGetLogin tests that GetLogin properly extracts login information from the Keeper record.
func TestGetLogin(t *testing.T) {
	// Mock values
	uid := "boop"
	title := "my-test-login"
	notes := "my-notes"
	login := "test-user"
	password := "test-password"
	url := "https://test.com"

	// Example JSON data that comes from the Keeper API for a login record
	loginRecordJson := fmt.Sprintf(`{
    "uid": "%s",
    "title": "%s",
    "type": "login",
    "notes": "%s",
    "fields": [
        {
            "label": "passkey",
            "type": "passkey",
            "value": []
        },
        {
            "label": "login",
            "type": "login",
            "value": [
                "%s"
            ]
        },
        {
            "label": "password",
            "type": "password",
            "value": [
                "%s"
            ]
        },
        {
            "label": "url",
            "type": "url",
            "value": [
                "%s"
            ]
        },
        {
            "label": "fileRef",
            "type": "fileRef",
            "value": []
        },
        {
            "label": "oneTimeCode",
            "type": "oneTimeCode",
            "value": []
        }
    ],
    "custom_fields": [],
    "files": []
}`,uid, title, notes, login, password, url)

    // Create a mocked client with the example JSON data. 
	// The client will return the JSON data as a Keeper record.
	client := getMockedClient(loginRecordJson)
	loginRecord, err := client.GetLogin(uid)

	// Assert that the returned record matches the expected values.
	require.NoError(t, err)
	assert.Equal(t, title, loginRecord.Title)
	assert.Equal(t, notes, loginRecord.Notes)
	assert.Equal(t, login, loginRecord.Login)
	assert.Equal(t, password, loginRecord.Password)
	assert.Equal(t, url, loginRecord.Url)
	assert.Equal(t, uid, loginRecord.Uid)
	assert.Equal(t, LOGIN_FIELD_TYPE, loginRecord.Type)
}

// TestGetAPIKey tests that GetApiKey properly extracts API key information from the Keeper record.
func TestGetApiKey(t *testing.T) {
	// Mock values
	uid := "54c1e2f3-8a0b-4d5e-8b7f-9c6d7e8f9a0b"
	title := "Test My API Key"
	notes := "cool-note"
	appID := "test-app-id"
	clientSecret := "12345"

	// Example JSON data that comes from the Keeper API for an API key record
	apiKeyJson := fmt.Sprintf(`{
    "uid": "%s",
    "title": "%s",
    "type": "API Key",
    "notes": "%s",
    "fields": [
        {
            "label": "AppID",
            "type": "text",
            "value": [
                "%s"
            ]
        },
        {
            "label": "ClientSecret",
            "type": "text",
            "value": [
                "%s"
            ]
        },
        {
            "label": "Cert",
            "type": "fileRef",
            "value": []
        }
    ],
    "custom_fields": [],
    "files": []
	
}`, uid, title, notes, appID, clientSecret)

	// Create a mocked client with the example JSON data.
	client := getMockedClient(apiKeyJson)
	apiKeyRecord, err := client.GetApiKey(uid)
	require.NoError(t, err)

	// Assert that the returned record matches the expected values.
	
	assert.Equal(t, title, apiKeyRecord.Title)
	assert.Equal(t, notes, apiKeyRecord.Notes)
	assert.Equal(t, appID, apiKeyRecord.AppId)
	assert.Equal(t, clientSecret, apiKeyRecord.ClientSecret)
	assert.Equal(t, uid, apiKeyRecord.Uid)
	assert.Equal(t, API_KEY_FIELD_TYPE, apiKeyRecord.Type)
}

// TestGetSoftwareLicense tests that GetSoftwareLicense properly extracts software license information from the Keeper record.
func TestGetSoftwareLicense(t *testing.T) {
	// Mock data for the license record
	uid := "5we3f4b2-8a0b-4d5e-8b7f-9c6d7e8f9a0b"
	title := "Test Software License"
	notes := "best license"
	licenseNumber := "12345"

	// We return the date as a string so we need to convert it to a string
	expirationDate := fmt.Sprintf("%d", time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli())
	activationDate := fmt.Sprintf("%d", time.Date(2025, 1, 2, 0, 0, 0, 0, time.UTC).UnixMilli())

	// Convert the date to the expected format YY-MM-DDTHH:MM:SS-07:00
	expectedActivationDate := ConvertDateStr(activationDate)
	expectedExpirationDate := ConvertDateStr(expirationDate)

	// Example JSON data that comes from the Keeper API for a software license record
	jsonData := fmt.Sprintf(`{
    "uid": "%s",
    "title": "%s",
    "type": "softwareLicense",
    "notes": "%s",
    "fields": [
        {
            "label": "licenseNumber",
            "type": "licenseNumber",
            "value": [
                "%s"
            ]
        },
        {
            "label": "expirationDate",
            "type": "expirationDate",
            "value": [
                "%s"
            ]
        },
        {
            "label": "dateActive",
            "type": "date",
            "value": [
                "%s"
            ]
        },
        {
            "label": "fileRef",
            "type": "fileRef",
            "value": []
        }
    ],
    "custom_fields": [],
    "files": []
}`, uid, title, notes, licenseNumber, expirationDate, activationDate)

	// Create a mocked client with the example JSON data.
	client := getMockedClient(jsonData)
	softwareLicenseRecord, err := client.GetSoftwareLicense(uid)
	require.NoError(t, err)

	// Assert that the returned record matches the expected values.
	assert.Equal(t, title, softwareLicenseRecord.Title)
	assert.Equal(t, notes, softwareLicenseRecord.Notes)
	assert.Equal(t, licenseNumber, softwareLicenseRecord.LicenseNumber)
	assert.Equal(t, uid, softwareLicenseRecord.Uid)
	assert.Equal(t, softwareLicenseRecord.ActivationDate, expectedActivationDate)
	assert.Equal(t, softwareLicenseRecord.ExpirationDate, expectedExpirationDate)
	assert.Equal(t, SOFTWARE_LICENSE_FIELD_TYPE, softwareLicenseRecord.Type)	
}

// TestGetFile tests that GetFile properly extracts file information from the Keeper record.
func TestGetFile(t *testing.T) {
	// Using a template to generate the JSON data for the file record
	type templateData struct{
		Uid string
		Title string
		Note string
		FileRefs []FileRef
	}
	
	// Mock data for the file record
	data := &templateData{
		Uid: "test-uid",
		Title: "test-title",
		Note: "test-note",
		FileRefs: []FileRef{
			{
				Uid: "file-uid-1",
				Name: "file-name-1",
				Title: "file-title-1",
				Type: "text/plain",
				Size: 1234,
				LastModified: int(time.Now().UnixMilli()),
			},
			{
				Uid: "file-uid-2",
				Name: "file-name-2",
				Title: "file-title-2",
				Type: "text/plain",
				Size: 5678,
				LastModified: int(time.Now().UnixMilli()),
			},
		},
	}

	// Template for the JSON data
	jsonData := fmt.Sprintf(`{
    "uid": "{{ .Uid }}",
    "title": "{{ .Title }}",
    "type": "file",
    "notes": "{{ .Note }}",
    "fields": [
        {
            "label": "fileRef",
            "type": "fileRef",
            "value": []
        }
    ],
    "custom_fields": [],
    "files": [
		{{- range $index, $file := .FileRefs }}
		{
			"uid": "{{ $file.Uid }}",
			"name": "{{ $file.Name }}",
			"title": "{{ $file.Title }}",
			"type": "{{ $file.Type }}",
			"last_modified": {{ $file.LastModified }},
			"size": {{ $file.Size }}
		}{{ if ne (add $index 1) (len $.FileRefs) }},{{ end }}
		{{- end }}
    ]
}`)

	// Add a custom add func to the template
	addFunc := func (a, b int) int {
		return a + b
	}

	// Create the template
	template, err := template.New("files").Funcs(template.FuncMap{"add": addFunc}).Parse(jsonData)
	require.NoError(t, err)

	// Execute the template with the data and write the data to a buffer
	bfrWriter := bytes.NewBuffer(nil)
	err = template.Execute(bfrWriter, data)
	require.NoError(t, err)

	// Create a mocked client with the generated JSON data
	client := getMockedClient(bfrWriter.String())
	fileRecord, err := client.GetFile(data.Uid)
	require.NoError(t, err)

	// Assert that the returned record matches the expected values
	assert.Equal(t, data.Title, fileRecord.Title)
	assert.Equal(t, data.Note, fileRecord.Notes)
	assert.Equal(t, data.Uid, fileRecord.Uid)
	assert.Equal(t, FILE_FIELD_TYPE, fileRecord.Type)
	require.Equal(t, len(data.FileRefs), len(fileRecord.FileRefs))

	// Check each file reference to make sure it is the same as the original
	// this simulates when there are multiple files in the record.
	for i, fileRef := range data.FileRefs {
		assert.Equal(t, fileRef.Uid, fileRecord.FileRefs[i].Uid)
		assert.Equal(t, fileRef.Name, fileRecord.FileRefs[i].Name)
		assert.Equal(t, fileRef.Title, fileRecord.FileRefs[i].Title)
		assert.Equal(t, fileRef.Type, fileRecord.FileRefs[i].Type)
		assert.Equal(t, fileRef.Size, fileRecord.FileRefs[i].Size)
		assert.Equal(t, fileRef.LastModified, fileRecord.FileRefs[i].LastModified)
	}
}

// TestGetEncryptedNote tests that GetEncryptedNote properly extracts encrypted note information from the Keeper record.
func TestGetEncryptedNote(t *testing.T) {
	// Mock values
	uid := "test-uid"
	title := "test-title"
	notes := "test-notes"
	secretNote := "super-secret-note"

	// Keeper returns dates as strings, so this tests that we are converting
	// the date to the correct format.
	date := 1
	expectedDate := ConvertDateStr(fmt.Sprintf("%d", date))

	// Example JSON data that comes from the Keeper API for an encrypted note record
	jsonData := fmt.Sprintf(`{
    "uid": "%s",
    "title": "%s",
    "type": "encryptedNotes",
    "notes": "%s",
    "fields": [
        {
            "label": "note",
            "type": "note",
            "value": [
                "%s"
            ]
        },
        {
            "label": "date",
            "type": "date",
            "value": [
                "%d"
            ]
        },
        {
            "label": "fileRef",
            "type": "fileRef",
            "value": []
        }
    ],
    "custom_fields": [],
    "files": []
}`, uid, title, notes, secretNote, date)

	// Create a mocked client with the example JSON data.
	client := getMockedClient(jsonData)
	encryptedNoteRecord, err := client.GetEncryptedNote(uid)
	require.NoError(t, err)

	// Assert that the returned record matches the expected values.
	assert.Equal(t, title, encryptedNoteRecord.Title)
	assert.Equal(t, notes, encryptedNoteRecord.Notes)
	assert.Equal(t, secretNote, encryptedNoteRecord.Note)
	assert.Equal(t, uid, encryptedNoteRecord.Uid)
	assert.Equal(t, ENCRYPTED_NOTE_FIELD_TYPE, encryptedNoteRecord.Type)
	assert.Equal(t, expectedDate, encryptedNoteRecord.Date)
}

// TestGetServerCredentials tests that GetServerCredentials properly extracts server credentials information from the Keeper record.
func TestGetServerCredentials(t *testing.T) {
	// Mock values
	uid := "test-uid"
	title := "test-title"
	notes := "test-notes"
	login := "test-login"
	password := "test-password"
	hostName := "test-hostname"
	port := 1234

	// Example JSON data that comes from the Keeper API for a server credentials record
	jsonData := fmt.Sprintf(`{
	"uid": "%s",
	"title": "%s",
	"type": "serverCredentials",
	"notes": "%s",
	"fields": [
		{
			"label": "login",
			"type": "login",
			"value": [
				"%s"
			]
		},
		{
			"label": "password",
			"type": "password",
			"value": [
				"%s"
			]
		},
		{
		"label": "host",
            "type": "host",
            "value": [
                {
                    "hostName": "%s",
                    "port": "%d"
                }
            ]
		}
	],
	"custom_fields": [],
	"files": []
}`, uid, title, notes, login, password, hostName, port)

	// Create a mocked client with the example JSON data.
	client := getMockedClient(jsonData)
	serverCredentialsRecord, err := client.GetServerCredentials(uid)
	require.NoError(t, err)

	// Assert that the returned record matches the expected values.
	assert.Equal(t, title, serverCredentialsRecord.Title)
	assert.Equal(t, notes, serverCredentialsRecord.Notes)
	assert.Equal(t, login, serverCredentialsRecord.Login)
	assert.Equal(t, password, serverCredentialsRecord.Password)
	assert.Equal(t, hostName, serverCredentialsRecord.HostConnection.HostName)
	assert.Equal(t, port, serverCredentialsRecord.HostConnection.Port)
	assert.Equal(t, uid, serverCredentialsRecord.Uid)
	assert.Equal(t, SERVER_FIELD_TYPE, serverCredentialsRecord.Type)
}

// TestGetDatabaseCredentials tests that GetDatabaseCredentials properly extracts database credentials information from the Keeper record.
func TestGetDatabaseCredentials(t *testing.T) {
	// Mock values
	uid := "test-uid"
	title := "test-title"
	notes := "test-notes"
	login := "test-login"
	password := "test-password"
	hostName := "test-hostname"
	port := 1234

	// Example JSON data that comes from the Keeper API for a database credentials record
	jsonData := fmt.Sprintf(`{
	"uid": "%s",
	"title": "%s",
	"type": "databaseCredentials",
	"notes": "%s",
	"fields": [
		{
			"label": "login",
			"type": "login",
			"value": [
				"%s"
			]
		},
		{
			"label": "password",
			"type": "password",
			"value": [
				"%s"
			]
		},
		{
			"label": "host",
			"type": "host",
			"value": [
				{
					"hostName": "%s",
					"port": "%d"
				}
			]
		}
	],
	"custom_fields": [],
	"files": []
}`, uid, title, notes, login, password, hostName, port)

	// Create a mocked client with the example JSON data.
	client := getMockedClient(jsonData)
	databaseCredentialsRecord, err := client.GetDatabaseCredentials(uid)
	require.NoError(t, err)

	// Assert that the returned record matches the expected values.
	assert.Equal(t, title, databaseCredentialsRecord.Title)
	assert.Equal(t, notes, databaseCredentialsRecord.Notes)
	assert.Equal(t, login, databaseCredentialsRecord.Login)
	assert.Equal(t, password, databaseCredentialsRecord.Password)
	assert.Equal(t, hostName, databaseCredentialsRecord.HostConnection.HostName)
	assert.Equal(t, port, databaseCredentialsRecord.HostConnection.Port)
	assert.Equal(t, uid, databaseCredentialsRecord.Uid)
	assert.Equal(t, DATABASE_FIELD_TYPE, databaseCredentialsRecord.Type)
}

// TestConvertDateStr tests the ConvertDateStr function to ensure it converts a date string to the correct format.
// It tests both a valid date string and an invalid date string. In Keeper dates are not guaranteed to be set and may be invalid.
// for this reason it makes sense to return an empty string for missing dates.
func TestConvertDateStr(t *testing.T) {
	// Test the function with a valid date string
	t.Run("Valid date string", func(t *testing.T) {
	dateStr := "1672531199000"
	expectedDate := time.Unix(1672531199, 0).Format(time.RFC3339)
	result := ConvertDateStr(dateStr)
	assert.Equal(t, expectedDate, result)
	})

	// Test the function with an invalid date string
	t.Run("invalid date string", func(t *testing.T) {
		dateStr := "invalid-date"
		result := ConvertDateStr(dateStr)
		assert.Empty(t, result)
	})	

}

func getMockedClient(secretRecordJson string) *PackerKeeperClient {
	mockClient := &MockKeeperClient{
		TestClient: &KSMClient{},
	}

	record := recordFromJSON(secretRecordJson)
	mockClient.On("GetSecret").Return(record, nil)

	return &PackerKeeperClient{
		KeeperClient: mockClient,
	}
}

func recordFromJSON(data string) *ksm.Record {
	recordFieldsDict := ksm.JsonToDict(data)
	record :=  &ksm.Record{
		RecordDict: recordFieldsDict,
	}

	// JSONToDict only builds the dictionary so we need to setup
	// mock some fields ourselves based off the JSON payload.
	if uid, ok := recordFieldsDict["uid"].(string); ok{
		record.Uid = uid
	} else {
		record.Uid = ""
	}

	// Mock the file payload. Keeper takes the files dictionary and adds it to
	// .Record.Files so this just mocks that behavior.
	if files, ok := recordFieldsDict["files"].([]interface{}); ok {
		ksmFiles := []*ksm.KeeperFile{}
		for _, file := range files {
			fileDict, ok := file.(map[string]interface{})
			if !ok {
				continue
			}
			
			// Build a keeper file object from the dictionary
			keeperFile := &ksm.KeeperFile{
				Uid:          fileDict["uid"].(string),
				Name:         fileDict["name"].(string),
				Title:        fileDict["title"].(string),
				Type:         fileDict["type"].(string),
				LastModified: int(fileDict["last_modified"].(float64)),
				Size:         int(fileDict["size"].(float64)),
				FileData: []byte("string"),
			}

			ksmFiles = append(ksmFiles, keeperFile)
		}

		record.Files = ksmFiles
	}

	// Return our built record.
	return record
}
