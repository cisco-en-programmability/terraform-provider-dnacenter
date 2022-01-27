
data "dnacenter_event_api_status" "example" {
  provider     = dnacenter
  execution_id = "string"
}

output "dnacenter_event_api_status_example" {
  value = data.dnacenter_event_api_status.example.item
}
