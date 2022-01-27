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

data "dna_global_credentials" "response" {
  provider            = dnacenter
  credential_sub_type = "CLI"
}

output "dna_global_credentials_response" {
  value       = data.dna_global_credentials.response
  description = "The dna_global_credentials data source's response"
}
