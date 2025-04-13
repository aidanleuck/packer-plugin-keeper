# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-software-license" "test" {
  # Single file record
  uid = "C-Avh6Ou1MKG9saFfjQaMA"
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
      "echo Title: ${data.keeper-software-license.test.title}",
      "echo Notes: ${data.keeper-software-license.test.notes}",
      "echo License Key: ${data.keeper-software-license.test.license_number}",
      "echo Activation Date: ${data.keeper-software-license.test.expiration_date}",
      "echo Expiration Date: ${data.keeper-software-license.test.activation_date}",
    ]
  }
}
