
resource "dnacenter_itsm_integration_events_retry" "example" {
  provider   = dnacenter
  parameters = ["string"]
}

output "dnacenter_itsm_integration_events_retry_example" {
  value = dnacenter_itsm_integration_events_retry.example
}