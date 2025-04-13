# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-database-credential" "test" {
  # Single file record
  uid = "yd4Z7jSWUww72Vzk7EnNdg"
}

source "null" "basic-example" {
  communicator = "none"
}

build {
  sources = [
    "source.null.basic-example"
  ]

  provisioner "shell-local" {
    inline = [
      "echo Title: ${data.keeper-database-credential.test.title}",
      "echo Notes: ${data.keeper-database-credential.test.notes}",
      "echo Login: ${data.keeper-database-credential.test.login}",
      "echo Password: ${data.keeper-database-credential.test.password}",
      "echo Host: ${data.keeper-database-credential.test.connection_details.host_name}",
      "echo Port: ${data.keeper-database-credential.test.connection_details.port}",
      "echo Type: ${data.keeper-database-credential.test.db_type}",
    ]
  }
}
