terraform {
  required_providers {
    dnacenter = {
      version = "1.3.1-beta"
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
  limit    = 1
  name     = "string"
  offset   = 1
}

output "dnacenter_nfv_profile_example" {
  value = data.dnacenter_nfv_profile.example.items
}
