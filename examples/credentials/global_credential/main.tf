terraform {
  required_providers {
    dnacenter = {
      versions = ["0.2"]
      source   = "hashicorp.com/edu/dnacenter"
    }
  }
}

provider "dnacenter" {
}

data "dna_global_credentials" "response" {
  provider            = dnacenter
  credential_sub_type = "CLI"
}

output "dna_global_credentials_response" {
  value       = data.dna_global_credentials.response
  description = "The dna_global_credentials data source's response"
}
