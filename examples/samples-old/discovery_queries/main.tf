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

data "dnacenter_discovery_range" "response" {
  provider          = dnacenter
  start_index       = 1
  records_to_return = 2
}

output "dnacenter_discovery_range_response" {
  value = data.dnacenter_discovery_range.response.items
}