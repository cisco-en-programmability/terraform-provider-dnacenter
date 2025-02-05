
data "dnacenter_icap_settings" "example" {
  provider       = dnacenter
  apid           = "string"
  capture_status = "string"
  capture_type   = "string"
  client_mac     = "string"
  limit          = 1
  offset         = 1
  wlc_id         = "string"
}

output "dnacenter_icap_settings_example" {
  value = data.dnacenter_icap_settings.example.items
}
