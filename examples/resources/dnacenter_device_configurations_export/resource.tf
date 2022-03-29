
resource "dnacenter_device_configurations_export" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    device_id = ["string"]
    password  = "******"
  }
}