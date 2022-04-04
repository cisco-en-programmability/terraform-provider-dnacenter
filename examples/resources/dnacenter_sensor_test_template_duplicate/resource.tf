
resource "dnacenter_sensor_test_template_duplicate" "example" {
  provider = dnacenter
  lifecycle {
    create_before_destroy = true
  }
  parameters {
    new_template_name = "string"
    template_name     = "string"
  }
}