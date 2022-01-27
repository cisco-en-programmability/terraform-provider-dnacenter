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

data "dnacenter_applications" "example" {
  provider = dnacenter
  limit    = 1
  name     = "intrinsa"
  offset   = 1
}

output "dnacenter_applications_example" {
  value = data.dnacenter_applications.example.items
}
