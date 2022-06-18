
resource "dnacenter_sensor_test_run" "example" {
  provider = dnacenter
  parameters {

    template_name = "string"
  }
}

output "dnacenter_sensor_test_run_example" {
  value = dnacenter_sensor_test_run.example
}