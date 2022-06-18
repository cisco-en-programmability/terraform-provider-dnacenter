
resource "dnacenter_pnp_device_config_preview" "example" {
  provider = dnacenter
  parameters {

    device_id = "string"
    site_id   = "string"
    type      = "string"
  }
}

output "dnacenter_pnp_device_config_preview_example" {
  value = dnacenter_pnp_device_config_preview.example
}