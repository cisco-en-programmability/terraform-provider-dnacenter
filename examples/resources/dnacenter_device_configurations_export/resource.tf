
resource "dnacenter_device_configurations_export" "example" {
  provider = dnacenter
  parameters {

    device_id = ["string"]
    password  = "******"
  }
}

output "dnacenter_device_configurations_export_example" {
  value = dnacenter_device_configurations_export.example
}