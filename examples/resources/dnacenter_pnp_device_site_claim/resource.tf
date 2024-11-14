
resource "dnacenter_pnp_device_site_claim" "example" {
  provider = dnacenter
  parameters {

    config_info {

      config_id = "string"
      config_parameters {

        key   = "string"
        value = "string"
      }
    }
    device_id = "string"
    gateway   = "string"
    hostname  = "string"
    image_info {

      image_id = "string"
      skip     = "false"
    }
    ip_interface_name = "string"
    rf_profile        = "string"
    sensor_profile    = "string"
    site_id           = "string"
    static_ip         = "string"
    subnet_mask       = "string"
    type              = "string"
    vlan_id           = "string"
  }
}

output "dnacenter_pnp_device_site_claim_example" {
  value = dnacenter_pnp_device_site_claim.example
}