
data "dnacenter_compliance_check_run" "example" {
  provider     = dnacenter
  categories   = ["string"]
  device_uuids = ["string"]
  trigger_full = "false"
}