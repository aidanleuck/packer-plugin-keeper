# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-file" "test" {
  # Single file record
  uid = "OIPg0wxSeC9wiuvTZH0N-w"
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
      "echo File UID: ${data.keeper-file.test.file_refs[0].uid}",
      "echo File Name: ${data.keeper-file.test.file_refs[0].name}",
      "echo Title: ${data.keeper-file.test.title}",
      "echo Size: ${data.keeper-file.test.file_refs[0].size}",
      "echo Last Modified: ${data.keeper-file.test.file_refs[0].last_modified}",
      "echo Base 64 Data: ${data.keeper-file.test.file_refs[0].content_base64}",
      "echo Type: ${data.keeper-file.test.file_refs[0].type}",
      "echo Notes: ${data.keeper-file.test.notes}",
    ]
  }
}
