terraform {
  required_providers {
    dnacenter = {
      version = "1.1.5-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

data "dnacenter_interface" "example" {
  provider       = dnacenter
  interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
}

output "dnacenter_interface_example" {
  value = data.dnacenter_interface.example.item
}
