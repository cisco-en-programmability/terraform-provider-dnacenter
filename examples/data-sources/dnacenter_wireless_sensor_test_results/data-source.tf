
data "dnacenter_wireless_sensor_test_results" "example" {
  provider        = dnacenter
  end_time        = "hh:mm"
  site_id         = "string"
  start_time      = "hh:mm"
  test_failure_by = "string"
}

output "dnacenter_wireless_sensor_test_results_example" {
  value = data.dnacenter_wireless_sensor_test_results.example.item
}
