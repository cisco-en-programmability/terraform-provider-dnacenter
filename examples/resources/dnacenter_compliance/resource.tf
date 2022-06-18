
resource "dnacenter_compliance" "example" {
  provider = dnacenter
  parameters {

    categories   = ["string"]
    device_uuids = ["string"]
    trigger_full = "false"
  }
}

output "dnacenter_compliance_example" {
  value = dnacenter_compliance.example
}