
data "dnacenter_pnp_device_claim_to_site" "example" {
  provider  = dnacenter
  device_id = "string"
  site_id   = "string"
  type      = "string"
}