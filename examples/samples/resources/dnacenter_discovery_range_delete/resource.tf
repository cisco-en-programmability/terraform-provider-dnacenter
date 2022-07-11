
terraform {
  required_providers {
    dnacenter = {
      version = "1.0.3-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_discovery_range_delete" "example" {
  provider = dnacenter
 
  parameters {
    records_to_delete = 1
    start_index       = 1
  }
}