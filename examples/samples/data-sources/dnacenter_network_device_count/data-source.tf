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

data "dnacenter_network_device_count" "example" {
  provider = dnacenter
  device_id = "24241ac1-75ad-4288-8284-145315182785"
}

output "dnacenter_network_device_count_example" {
  value = data.dnacenter_network_device_count.example.item
}
