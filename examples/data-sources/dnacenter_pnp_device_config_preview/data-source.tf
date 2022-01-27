
data "dnacenter_pnp_device_config_preview" "example" {
  provider  = dnacenter
  device_id = "string"
  site_id   = "string"
  type      = "string"
}