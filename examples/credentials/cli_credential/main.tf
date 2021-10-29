terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source   = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
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
  sensitive = true
  value = dna_cli_credential.response
}
