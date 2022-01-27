
terraform {
  required_providers {
  dnacenter = {
    version = "0.0.3"
    source  = "hashicorp.com/edu/dnacenter"
    # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
  }
  }
}
resource "dnacenter_network_device" "example" {
  provider = dnacenter
  parameters {

    id = "3eb928b8-2414-4121-ac35-1247e5d666a4"
  }
}

output "dnacenter_network_device_example" {
  value = dnacenter_network_device.example
}