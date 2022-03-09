
resource "dnacenter_pnp_device_site_claim" "example" {
  provider = dnacenter
  parameters {
    device_id = "string"
    site_id   = "string"
    type      = "string"
    hostname  = "string"
    image_info {
      image_id = "string"
      skip     = "false"
    }
    config_info {
      config_id = "string"
      config_parameters {
        key   = "string"
        value = "string"
      }
    }
  }
}

output "dnacenter_pnp_device_site_claim_example" {
  value = dnacenter_pnp_device_site_claim.example
}
