
resource "dnacenter_telemetry_settings_apply" "example" {
  provider = dnacenter
  parameters {

    device_ids = ["string"]
  }
}

output "dnacenter_telemetry_settings_apply_example" {
  value = dnacenter_telemetry_settings_apply.example
}