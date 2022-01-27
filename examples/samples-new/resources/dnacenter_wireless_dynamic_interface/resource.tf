terraform {
  required_providers {
    dnacenter = {
      version = "0.0.3"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source, change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
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