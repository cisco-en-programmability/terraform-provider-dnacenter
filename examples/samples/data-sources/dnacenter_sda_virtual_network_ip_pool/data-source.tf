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

data "dnacenter_sda_virtual_network_ip_pool" "example" {
  provider             = dnacenter
  ip_pool_name         = "string"
  virtual_network_name = "string"
}

output "dnacenter_sda_virtual_network_ip_pool_example" {
  value = data.dnacenter_sda_virtual_network_ip_pool.example.item
}
