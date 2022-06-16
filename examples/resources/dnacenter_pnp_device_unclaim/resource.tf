
resource "dnacenter_pnp_device_unclaim" "example" {
  provider = dnacenter
  parameters {

    device_id_list = ["string"]
  }
}

output "dnacenter_pnp_device_unclaim_example" {
  value = dnacenter_pnp_device_unclaim.example
}