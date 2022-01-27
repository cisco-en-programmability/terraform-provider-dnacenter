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

data "dnacenter_sda_port_assignment_for_user_device" "example" {
  provider                     = dnacenter
  device_management_ip_address = "string"
  interface_name               = "string"
}

output "dnacenter_sda_port_assignment_for_user_device_example" {
  value = data.dnacenter_sda_port_assignment_for_user_device.example.item
}
