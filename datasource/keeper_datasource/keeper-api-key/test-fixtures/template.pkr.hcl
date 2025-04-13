# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-api-key" "test" {
  # Test API Key Record
  uid = "2SbJe2lVPycVrK_0Vmnw-g"
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
      "echo Title: ${data.keeper-api-key.test.title}",
      "echo ClientSecret: ${data.keeper-api-key.test.client_secret}",
      "echo AppID: ${data.keeper-api-key.test.app_id}",
      "echo Notes: ${data.keeper-api-key.test.notes}",
    ]
  }
}
