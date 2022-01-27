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

data "dnacenter_network_device_count" "example" {
  provider  = dnacenter
  device_id = "string"
}

output "dnacenter_network_device_count_example" {
  value = data.dnacenter_network_device_count.example.item_name
}

output "dnacenter_network_device_count_example" {
  value = data.dnacenter_network_device_count.example.item_id
}
