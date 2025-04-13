# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-encrypted-note" "test" {
  # Single file record
  uid = "MFs38NixUjlA4qLJEoC-9w"
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
      "echo Title: ${data.keeper-encrypted-note.test.title}",
      "echo notes: ${data.keeper-encrypted-note.test.notes}",
      "echo securedNote: ${data.keeper-encrypted-note.test.note}",
      "echo date: ${data.keeper-encrypted-note.test.date}",
    ]
  }
}
