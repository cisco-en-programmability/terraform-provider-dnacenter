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

data "dnacenter_sda_port_assignment_for_access_point" "example" {
  provider                     = dnacenter
  device_management_ip_address = "10.121.1.1"
  interface_name               = "string"
}

output "dnacenter_sda_port_assignment_for_access_point_example" {
  value = data.dnacenter_sda_port_assignment_for_access_point.example.item
}
