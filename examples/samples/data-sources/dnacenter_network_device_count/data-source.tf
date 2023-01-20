terraform {
  required_providers {
    dnacenter = {
      version = "1.0.17-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
}

data "dnacenter_network_device_count" "example" {
  provider = dnacenter
  #device_id = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
}

output "dnacenter_network_device_count_example" {
  value = data.dnacenter_network_device_count.example.item
}
