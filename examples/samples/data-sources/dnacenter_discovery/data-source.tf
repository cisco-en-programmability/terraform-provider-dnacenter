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

data "dnacenter_discovery" "example" {
  provider = dnacenter
  id       = "6"
}

output "dnacenter_discovery_example" {
  value = data.dnacenter_discovery.example.item
}
