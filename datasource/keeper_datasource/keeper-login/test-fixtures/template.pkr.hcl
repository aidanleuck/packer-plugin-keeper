# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

data "keeper-login" "test" {
  # Test login record
  uid = "A6En9kNc6HppPWDOi3MH9g"
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
      "echo Title: ${data.keeper-login.test.title}",
      "echo Url: ${data.keeper-login.test.url}",
      "echo Login: ${data.keeper-login.test.login}",
      "echo Password: ${data.keeper-login.test.password}",
      "echo Notes: ${data.keeper-login.test.notes}",
    ]
  }
}
