
data "dnacenter_wireless_access_points_factory_reset_request_status" "example" {
  provider = dnacenter
  task_id  = "string"
}

output "dnacenter_wireless_access_points_factory_reset_request_status_example" {
  value = data.dnacenter_wireless_access_points_factory_reset_request_status.example.items
}
