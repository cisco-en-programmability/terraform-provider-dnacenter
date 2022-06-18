
data "dnacenter_wireless_sensor_test_results" "example" {
  provider        = dnacenter
  end_time        = 1609459200
  site_id         = "string"
  start_time      = 1609459200
  test_failure_by = "string"
}

output "dnacenter_wireless_sensor_test_results_example" {
  value = data.dnacenter_wireless_sensor_test_results.example.item
}
