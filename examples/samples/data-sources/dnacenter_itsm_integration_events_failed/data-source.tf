
data "dnacenter_itsm_integration_events_failed" "example" {
  provider    = dnacenter
  instance_id = "string"
}

output "dnacenter_itsm_integration_events_failed_example" {
  value = data.dnacenter_itsm_integration_events_failed.example.items
}
