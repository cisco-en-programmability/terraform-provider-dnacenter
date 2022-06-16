
terraform {
  required_providers {
    dnacenter = {
      version = "0.3.0"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}

provider "dnacenter" {
  debug = "true"
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

resource "dnacenter_compliance" "example" {
  depends_on = [dnacenter_network_device.example]
  provider   = dnacenter
  parameters {
    trigger_full = true
    categories   = ["PSIRT"]
    device_uuids = [dnacenter_network_device.example.item.0.id]
  }
}

output "dnacenter_compliance_example" {
  value = dnacenter_compliance.example
}
