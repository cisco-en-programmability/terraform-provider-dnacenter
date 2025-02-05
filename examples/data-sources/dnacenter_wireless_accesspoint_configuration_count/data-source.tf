
data "dnacenter_wireless_accesspoint_configuration_count" "example" {
  provider       = dnacenter
  ap_mode        = "string"
  ap_model       = "string"
  mesh_role      = "string"
  provisioned    = "string"
  wlc_ip_address = "string"
}

output "dnacenter_wireless_accesspoint_configuration_count_example" {
  value = data.dnacenter_wireless_accesspoint_configuration_count.example.item
}
