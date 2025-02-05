
data "dnacenter_security_advisories_results_network_devices_network_device_id_advisories" "example" {
  provider               = dnacenter
  cvss_base_score        = "string"
  id                     = "string"
  limit                  = 1
  network_device_id      = "string"
  offset                 = 1
  order                  = "string"
  security_impact_rating = "string"
  sort_by                = "string"
}

output "dnacenter_security_advisories_results_network_devices_network_device_id_advisories_example" {
  value = data.dnacenter_security_advisories_results_network_devices_network_device_id_advisories.example.items
}
