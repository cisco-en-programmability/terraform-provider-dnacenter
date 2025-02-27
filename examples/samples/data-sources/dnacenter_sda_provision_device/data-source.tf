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

data "dnacenter_sda_provision_device" "example" {
  provider                     = dnacenter
  device_management_ip_address = "10.121.1.1"
}

output "dnacenter_sda_provision_device_example" {
  value = data.dnacenter_sda_provision_device.example.item
}
