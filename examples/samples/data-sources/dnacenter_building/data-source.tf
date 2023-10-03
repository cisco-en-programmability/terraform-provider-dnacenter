terraform {
  required_providers {
    dnacenter = {
      version = "1.1.19-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
data "dnacenter_building" "example" {
  provider = dnacenter
  # limit    = 1
  # name     = "string"
  # offset   = 1
  # site_id  = "string"
  # type     = "string"
}

output "dnacenter_building_example" {
  value = data.dnacenter_building.example.items
}
