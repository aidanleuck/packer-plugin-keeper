# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

packer {
  required_plugins {
    name = {
      source  = "github.com/aidanleuck/keeper"
      version = ">=0.0.1"
    }
  }
}

// Retrieve a login record
data "keeper-login" "my_login" {
  uid = "my-uid"
}

// Retrieve a ssh key record
data "keeper-ssh-key" "my_ssh_key" {
  uid = "my-uid"
}

// Retrieve a software license record
data "keeper-software-license" "my_software_license" {
  uid = "my-uid"
}

// Retrieve a api key record
data "keeper-api-key" "my_api_key" {
  uid = "my-uid"
}

// Retrieve a server credential record
data "keeper-server-credential" "my_server_credential" {
  uid = "my-uid"
}

// Retrieve a file record
data "keeper-file" "my_file" {
  uid = "my-uid"
}

// Retrieve a database credential record
data "keeper-database-credential" "my_database_credential" {
  uid = "my-uid"
}

// Retieve a encrypted note record
data "keeper-encrypted-note" "my_encrypted_note" {
  uid = "my-uid"
}
