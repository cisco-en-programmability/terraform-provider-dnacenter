
data "dnacenter_wireless_accesspoint_configuration_summary" "example" {
  provider       = dnacenter
  ap_mode        = "string"
  ap_model       = "string"
  key            = "string"
  limit          = 1
  mesh_role      = "string"
  offset         = 1
  provisioned    = "string"
  wlc_ip_address = "string"
}

output "dnacenter_wireless_accesspoint_configuration_summary_example" {
  value = data.dnacenter_wireless_accesspoint_configuration_summary.example.item
}
