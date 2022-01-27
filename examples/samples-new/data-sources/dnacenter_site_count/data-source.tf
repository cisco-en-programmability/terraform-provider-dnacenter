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

data "dnacenter_site_count" "example" {
  provider = dnacenter
  site_id  = "3e0db2cd-cf3a-4dbd-bfb9-739271ffc20b"
}

output "dnacenter_site_count_example" {
  value = data.dnacenter_site_count.example.item
}
