
resource "dnacenter_sensor_test_delete" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    template_name = "string"
  }
}