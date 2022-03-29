
resource "dnacenter_compliance_check_run" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    categories   = ["string"]
    device_uuids = ["string"]
    trigger_full = "false"
  }
}