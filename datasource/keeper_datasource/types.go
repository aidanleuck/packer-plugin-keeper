//go:generate packer-sdc struct-markdown
//go:generate packer-sdc mapstructure-to-hcl2 -type KeeperLogin,FileRef,KeeperEncryptedNote,KeeperFile,KeeperRecordField,KeeperSoftwareLicense,KeeperSSHKey,KeyPair,HostConnection,KeeperServerCredentials,KeeperDataBaseCredentials,Config

package keeper_datasource

type KeeperRecordField struct {
	// uid is the unique identifier for the record .
	Uid string `mapstructure:"uid"`
	// type is the type of the record . (ex: login, file, etc.)
	Type string `mapstructure:"type"`
	// title is the title or name of the record .
	Title string `mapstructure:"title"`
	// notes are the notes associated with the record .
	Notes string `mapstructure:"notes"`
	// FileRefs contain the list of file references associated with a record. See [FileRef](#nested-schema-for-fileref)
	FileRefs []FileRef `mapstructure:"file_refs"`
}

type KeeperLogin struct {
	KeeperRecordField `mapstructure:",squash"`
	// login is the username or email address of the record.
	Login string `mapstructure:"login"`
	// password is the password of the record.
	Password string `mapstructure:"password"`
	// url is the url associated withthe record.
	Url string `mapstructure:"url"`
}

type FileRef struct {
	// uid is the unique identifier for the file .
	Uid string `mapstructure:"uid"`
	// title is the title or name of the file .
	Title string `mapstructure:"title"`
	// name is the name of the file .
	Name string `mapstructure:"name"`
	// type is the type of the file .
	Type string `mapstructure:"type"`
	// size is the size of the file .
	Size int `mapstructure:"size"`
	// last_modified is the last modified date of the file .
	LastModified int `mapstructure:"last_modified"`
	// content_base64 is the base64 encoded content of the file .
	Base64Data string `mapstructure:"content_base64"`
}

type KeeperEncryptedNote struct {
	KeeperRecordField `mapstructure:",squash"`
	// note is the secret note content.
	Note string `mapstructure:"note"`
	// date is the date associated with the note.
	Date string `mapstructure:"date"`
}

type KeeperFile struct {
	KeeperRecordField `mapstructure:",squash"`
}

type KeeperSoftwareLicense struct {
	KeeperRecordField `mapstructure:",squash"`
	// license_number is the license number associated with the software.
	LicenseNumber string `mapstructure:"license_number"`
	// activation_date is the activation date of the software.
	ActivationDate string `mapstructure:"activation_date"`
	// expiration_date is the expiration date of the software.
	ExpirationDate string `mapstructure:"expiration_date"`
}

type KeeperSSHKey struct {
	KeeperRecordField `mapstructure:",squash"`
	Login             string         `mapstructure:"login"`
	Passphrase        string         `mapstructure:"passphrase"`
	KeyPair           KeyPair        `mapstructure:"key_pair"`
	HostConnection    HostConnection `mapstructure:"connection_details"`
}

type KeyPair struct {
	// Fixed: Changed mapstructure tag from "passphrase" to "public_key" to match field name
	PublicKey  string `mapstructure:"public_key"`
	PrivateKey string `mapstructure:"private_key"`
}

type HostConnection struct {
	// host_name is the name of the host to connect to.
	HostName string `mapstructure:"host_name"`
	// port is the port to connect to.
	Port int `mapstructure:"port"`
}

type KeeperServerCredentials struct {
	KeeperRecordField `mapstructure:",squash"`
	// connection_details are the connection details to connect to the server.
	// see [HostConnection](#nested-schema-for-hostconnection)
	HostConnection HostConnection `mapstructure:"connection_details"`
	// login is the username used to connect to the server.
	Login string `mapstructure:"login"`
	// password is the password used to connect to the server.
	Password string `mapstructure:"password"`
}

type KeeperDataBaseCredentials struct {
	KeeperRecordField `mapstructure:",squash"`
	// connection_details are the connection details to connect to the server.
	// see [HostConnection](#nested-schema-for-hostconnection)
	HostConnection HostConnection `mapstructure:"connection_details"`
	// login is the username used to connect to the server.
	Login string `mapstructure:"login"`
	// password is the password used to connect to the server.
	Password string `mapstructure:"password"`
	// db_type is the type of the database (ex: mysql, postgres, etc.)
	// it can also be used as the name of the database to connect to.
	DbType string `mapstructure:"db_type"`
}

type KeeperAPIKey struct {
	KeeperRecordField `mapstructure:",squash"`
	// app_id is the application id associated with the API key.
	AppId string `mapstructure:"app_id"`
	// client_secret is the secret associated with the API key.
	ClientSecret string `mapstructure:"client_secret"`
}

type Config struct {
	// Uid is the unique identifier for the record .
	// required `true`
	Uid *string `mapstructure:"uid" required:"true"`
}
