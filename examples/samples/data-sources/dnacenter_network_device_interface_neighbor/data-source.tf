terraform {
  required_providers {
    dnacenter = {
      version = "1.0.19-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}


data "dnacenter_network_device_interface_neighbor" "example" {
  provider       = dnacenter
  device_uuid    = "3923aed0-16e5-4ed0-b430-ff6dcfd9c517"
  interface_uuid = "c6820b57-ecde-4b6d-98db-06ba10486809"
}

output "dnacenter_network_device_interface_neighbor_example" {
  value = data.dnacenter_network_device_interface_neighbor.example.item
}
