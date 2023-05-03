terraform {
  required_providers {
    dnacenter = {
      version = "1.1.2-beta"
      source  = "hashicorp.com/edu/dnacenter"
      # "hashicorp.com/edu/dnacenter" is the local built source change to "cisco-en-programmability/dnacenter" to use downloaded version from registry
    }
  }
}
resource "dnacenter_pnp_device_authorize" "example" {
  provider = dnacenter
  parameters {

    device_id_list = ["618c2b7ff6b3b66f71acbaf8"]
  }
}

output "dnacenter_pnp_device_authorize_example" {
  value = dnacenter_pnp_device_authorize.example
}