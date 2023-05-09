terraform {
  required_providers {
    dnacenter = {
      version = "1.1.6-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_system_health_count" "example" {
  provider  = dnacenter
  domain    = "string"
  subdomain = "string"
}

output "dnacenter_system_health_count_example" {
  value = data.dnacenter_system_health_count.example.item
}
