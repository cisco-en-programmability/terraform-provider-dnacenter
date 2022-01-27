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

data "dnacenter_network_device_module_count" "example" {
  provider  = dnacenter
  device_id = "3eb928b8-2414-4121-ac35-1247e5d666a4"
}

output "dnacenter_network_device_module_count_example" {
  value = data.dnacenter_network_device_module_count.example.item
}
