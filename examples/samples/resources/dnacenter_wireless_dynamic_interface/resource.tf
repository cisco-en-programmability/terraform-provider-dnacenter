terraform {
  required_providers {
    dnacenter = {
      version = "1.1.22-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
}

resource "dnacenter_wireless_dynamic_interface" "example" {
  provider = dnacenter
  parameters {

    interface_name = "hello"
    //vlan_id        = 1.0
  }
}

output "dnacenter_wireless_dynamic_interface_example" {
  value = dnacenter_wireless_dynamic_interface.example
}
