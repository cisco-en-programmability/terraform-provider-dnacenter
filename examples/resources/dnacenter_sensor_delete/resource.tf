
resource "dnacenter_sensor_delete" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    template_name = "string"
  }
}