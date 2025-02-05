
data "dnacenter_security_advisories_results_advisories_id_network_devices" "example" {
  provider          = dnacenter
  id                = "string"
  limit             = 1
  network_device_id = "string"
  offset            = 1
  order             = "string"
  scan_mode         = "string"
  scan_status       = "string"
  sort_by           = "string"
}

output "dnacenter_security_advisories_results_advisories_id_network_devices_example" {
  value = data.dnacenter_security_advisories_results_advisories_id_network_devices.example.items
}
