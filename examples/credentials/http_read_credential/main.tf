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

resource "dna_http_read_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    port            = 23
  }
}
output "dna_http_read_credential_response" {
  value = dna_http_read_credential.response
}
