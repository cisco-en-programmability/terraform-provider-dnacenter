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

data "dnacenter_global_pool" "example" {
  provider = dnacenter
  limit    = "1"
  offset   = "1"
}

output "dnacenter_global_pool_example" {
  value = data.dnacenter_global_pool.example.items
}
