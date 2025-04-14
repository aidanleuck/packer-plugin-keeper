# Keeper Database Datasource

Type: `keeper-database-credential`

This datasource retrieves a keeper database record and outputs its contents as HCL structures for use in your Packer templates.

## Examples

- Basic examples are available in the [examples](https://github.com/aidanleuck/packer-plugin-keeper/tree/main/example)
  directory of the GitHub repository.

## Configuration Reference

### Inputs

#### Required

<!-- Code generated from the comments of the Config struct in datasource/keeper_datasource/types.go; DO NOT EDIT MANUALLY -->

- `uid` (\*string) - Uid is the unique identifier for the record .
  required `true`

<!-- End of code generated from the comments of the Config struct in datasource/keeper_datasource/types.go; -->


### Outputs

<!-- Code generated from the comments of the KeeperRecordField struct in datasource/keeper_datasource/types.go; DO NOT EDIT MANUALLY -->

- `uid` (string) - uid is the unique identifier for the record .

- `type` (string) - type is the type of the record . (ex: login, file, etc.)

- `title` (string) - title is the title or name of the record .

- `notes` (string) - notes are the notes associated with the record .

- `file_refs` ([]FileRef) - FileRefs contain the list of file references associated with a record. See [FileRef](#nested-schema-for-fileref)

<!-- End of code generated from the comments of the KeeperRecordField struct in datasource/keeper_datasource/types.go; -->

<!-- Code generated from the comments of the KeeperDataBaseCredentials struct in datasource/keeper_datasource/types.go; DO NOT EDIT MANUALLY -->

- `connection_details` (HostConnection) - connection_details are the connection details to connect to the server.
  see [HostConnection](#nested-schema-for-hostconnection)

- `login` (string) - login is the username used to connect to the server.

- `password` (string) - password is the password used to connect to the server.

- `db_type` (string) - db_type is the type of the database (ex: mysql, postgres, etc.)
  it can also be used as the name of the database to connect to.

<!-- End of code generated from the comments of the KeeperDataBaseCredentials struct in datasource/keeper_datasource/types.go; -->


#### Nested Schema for HostConnection

<!-- Code generated from the comments of the HostConnection struct in datasource/keeper_datasource/types.go; DO NOT EDIT MANUALLY -->

- `host_name` (string) - host_name is the name of the host to connect to.

- `port` (int) - port is the port to connect to.

<!-- End of code generated from the comments of the HostConnection struct in datasource/keeper_datasource/types.go; -->


#### Nested Schema for FileRef

<!-- Code generated from the comments of the FileRef struct in datasource/keeper_datasource/types.go; DO NOT EDIT MANUALLY -->

- `uid` (string) - uid is the unique identifier for the file .

- `title` (string) - title is the title or name of the file .

- `name` (string) - name is the name of the file .

- `type` (string) - type is the type of the file .

- `size` (int) - size is the size of the file .

- `last_modified` (int) - last_modified is the last modified date of the file .

- `content_base64` (string) - content_base64 is the base64 encoded content of the file .

<!-- End of code generated from the comments of the FileRef struct in datasource/keeper_datasource/types.go; -->
