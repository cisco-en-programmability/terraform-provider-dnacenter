
data "dnacenter_pnp_device_claim_to_site" "example" {
  provider  = dnacenter
  device_id = "string"
  site_id   = "string"
  type      = "string"
  hostname  = "string"
  config_info{
    config_id= "string"
    config_parameters{
      key="string"
      value= "string"
    }
  }
  image_info{
    image_id= "string"
    skip= "false"
  }
}
