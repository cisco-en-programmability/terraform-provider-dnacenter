
resource "dnacenter_pnp_device_authorize" "example" {
  provider = dnacenter
  parameters {

    device_id_list = ["8b720cdd-d9f9-4e12-bdda-4cb511f1c047"]
  }
}

output "dnacenter_pnp_device_authorize_example" {
  value = dnacenter_pnp_device_authorize.example
}