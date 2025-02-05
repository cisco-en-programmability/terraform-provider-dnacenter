
data "dnacenter_icap_spectrum_sensor_reports" "example" {
  provider        = dnacenter
  ap_mac          = "string"
  data_type       = 1.0
  end_time        = 1609459200
  limit           = 1
  offset          = 1
  start_time      = 1609459200
  time_sort_order = "string"
  xca_lle_rid     = "string"
}

output "dnacenter_icap_spectrum_sensor_reports_example" {
  value = data.dnacenter_icap_spectrum_sensor_reports.example.items
}
