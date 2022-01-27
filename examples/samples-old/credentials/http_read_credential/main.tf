terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
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
  sensitive = true
  value     = dna_http_read_credential.response
}
