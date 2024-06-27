terraform {
  required_providers {
    dnacenter = {
      version = "1.1.33-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
data "dnacenter_site" "example" {
  provider = dnacenter
  # limit    = 1
  # name     = "string"
  # offset   = 1
  # site_id  = "string"
  # type     = "floor"
}

output "dnacenter_site_example" {
  value = data.dnacenter_site.example.items
}
