
resource "dnacenter_pnp_device_claim_to_site" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_id = "string"
    site_id   = "string"
    type      = "string"
    hostname  = "string"
    config_info {
      config_id = "string"
      config_parameters {
        key   = "string"
        value = "string"
      }
    }
    image_info {
      image_id = "string"
      skip     = "false"
    }
  }
}
