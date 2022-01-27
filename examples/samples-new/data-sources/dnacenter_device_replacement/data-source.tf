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

data "dnacenter_device_replacement" "example" {
  provider                         = dnacenter
  family                           = ["string"]
  faulty_device_name               = "string"
  faulty_device_platform           = "string"
  faulty_device_serial_number      = "string"
  limit                            = 1
  offset                           = 1
  replacement_device_platform      = "string"
  replacement_device_serial_number = "string"
  replacement_status               = ["string"]
  sort_by                          = "string"
  sort_order                       = "string"
}

output "dnacenter_device_replacement_example" {
  value = data.dnacenter_device_replacement.example.items
}
