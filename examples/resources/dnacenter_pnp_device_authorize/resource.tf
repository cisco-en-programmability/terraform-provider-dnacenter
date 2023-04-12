
resource "dnacenter_pnp_device_authorize" "example" {
  provider = dnacenter
  parameters {

    device_id_list = ["string"]
  }
}

output "dnacenter_pnp_device_authorize_example" {
  value = dnacenter_pnp_device_authorize.example
}