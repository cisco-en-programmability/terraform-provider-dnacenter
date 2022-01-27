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

data "dnacenter_application_sets" "example" {
  provider = dnacenter
  limit    = 1
  name     = "consumer-misc"
  offset   = 1
}

output "dnacenter_application_sets_example" {
  value = data.dnacenter_application_sets.example.items
}
