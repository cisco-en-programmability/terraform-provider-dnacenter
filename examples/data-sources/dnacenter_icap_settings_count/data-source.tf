
data "dnacenter_icap_settings_count" "example" {
  provider       = dnacenter
  apid           = "string"
  capture_status = "string"
  capture_type   = "string"
  client_mac     = "string"
  wlc_id         = "string"
}

output "dnacenter_icap_settings_count_example" {
  value = data.dnacenter_icap_settings_count.example.item
}
