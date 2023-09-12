terraform {
  required_providers {
    dnacenter = {
      version = "1.1.16-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
data "dnacenter_floor" "example" {
  provider = dnacenter
  # limit    = 1
  # name     = "string"
  # offset   = 1
  # site_id  = "string"
  # type     = "string"
}

output "dnacenter_floor_example" {
  value = data.dnacenter_floor.example.items
}
