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

data "dnacenter_nfv_profile" "example" {
  provider = dnacenter
  id       = "string"
  limit    = "string"
  name     = "string"
  offset   = "string"
}

output "dnacenter_nfv_profile_example" {
  value = data.dnacenter_nfv_profile.example.items
}
