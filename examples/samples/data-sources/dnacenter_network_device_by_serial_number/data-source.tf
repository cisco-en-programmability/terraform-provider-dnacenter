
terraform {
  required_providers {
    dnacenter = {
      version = "0.3.0"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_network_device_by_serial_number" "example" {
  provider      = dnacenter
  serial_number = "FOC2214Z084"
}

output "dnacenter_network_device_by_serial_number_example" {
  value = data.dnacenter_network_device_by_serial_number.example.item
}
