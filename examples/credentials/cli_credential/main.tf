terraform {
  required_providers {
    dnacenter = {
      versions = ["0.0.3"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

resource "dna_cli_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "123456"
    credential_type = "APP"
    id              = "6dcc1da3-3bbe-4d26-9b39-6efc38f6ae2e"
  }
}
output "dna_cli_credential_response" {
  value = dna_cli_credential.response
}
