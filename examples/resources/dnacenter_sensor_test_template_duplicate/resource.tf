
resource "dnacenter_sensor_test_template_duplicate" "example" {
  provider = dnacenter
  parameters {

    new_template_name = "string"
    template_name     = "string"
  }
}

output "dnacenter_sensor_test_template_duplicate_example" {
  value = dnacenter_sensor_test_template_duplicate.example
}