
resource "dnacenter_itsm_integration_events_retry" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    payload = ["string"]
  }
}