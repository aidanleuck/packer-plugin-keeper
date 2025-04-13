# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-server-credential" "test" {
  # Single file record
  uid = "PomuiS8aaB20UoHxKZ7hRg"
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
      "echo Title: ${data.keeper-server-credential.test.title}",
      "echo Notes: ${data.keeper-server-credential.test.notes}",
      "echo Login: ${data.keeper-server-credential.test.login}",
      "echo Password: ${data.keeper-server-credential.test.password}",
      "echo Host: ${data.keeper-server-credential.test.connection_details.host_name}",
      "echo Port: ${data.keeper-server-credential.test.connection_details.port}"
    ]
  }
}
