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

resource "dna_http_write_credential" "response" {
  provider = dnacenter
  item {
    username        = "usuario2"
    password        = "ThisI5aP4s_sW0rd"
    credential_type = "APP"
    port            = 24
    id              = "5abaa9c9-4470-46c4-90d6-107594164845"
  }
}
output "dna_http_write_credential_response" {
  value = dna_http_write_credential.response
}
