# Keeper Login Datasource

Type: `keeper-login`

This datasource retrieves a keeper login record and outputs its contents as HCL structures for use in your Packer templates.

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

<!-- Code generated from the comments of the KeeperLogin struct in datasource/keeper_datasource/types.go; DO NOT EDIT MANUALLY -->

- `login` (string) - login is the username or email address of the record.

- `password` (string) - password is the password of the record.

- `url` (string) - url is the url associated withthe record.

<!-- End of code generated from the comments of the KeeperLogin struct in datasource/keeper_datasource/types.go; -->


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
