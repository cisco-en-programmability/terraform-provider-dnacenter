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

data "dnacenter_reserve_ip_subpool" "example" {
  provider = dnacenter
  limit    = "2"
  offset   = "1"
  site_id  = "9e860d9e-6499-40d1-9645-4b45bd684219"
}

output "dnacenter_reserve_ip_subpool_example" {
  value = data.dnacenter_reserve_ip_subpool.example.items
}
